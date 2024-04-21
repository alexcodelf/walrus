package vcsrepo

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/seal-io/utils/stringx"
	"golang.org/x/exp/maps"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/vcs"
)

type Fetcher struct{}

func New() *Fetcher {
	return &Fetcher{}
}

// Fetch fills the template status.
func (l *Fetcher) Fetch(ctx context.Context, tmpl *walruscore.Template) (*walruscore.Template, error) {
	tempDir := filepath.Join(os.TempDir(), "seal-template-"+stringx.RandomHex(10))
	defer os.RemoveAll(tempDir)

	source := tmpl.Spec.VCSRepository

	// Clone.
	{
		tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
		if err != nil {
			return nil, err
		}

		opts := vcs.GitCloneOptions{
			URL:             source.URL,
			InsecureSkipTLS: !tlsVerify,
		}

		cloneCtx, cancel := context.WithTimeout(ctx, 10*time.Minute)
		defer cancel()

		_, err = vcs.GitClone(cloneCtx, tempDir, opts)
		if err != nil {
			return nil, err
		}
	}

	r, err := git.PlainOpen(tempDir)
	if err != nil {
		return nil, err
	}

	// Get URL.
	tmpl.Status.URL = source.URL

	// Get icon.
	{
		icon, err := gitRepoIconURL(r, source.URL)
		if err != nil {
			return nil, err
		}
		tmpl.Status.Icon = icon
	}

	// Get versions.
	{
		vers, versionSchema, err := getVersions(tmpl, r)
		if err != nil {
			return nil, err
		}
		tmplVers, err := genTemplateVersions(ctx, tmpl, vers, versionSchema)
		if err != nil {
			return nil, err
		}
		// Index remote versions.
		tmplVersReverseIndexer := make(map[string]int)
		for i, v := range tmplVers {
			tmplVersReverseIndexer[v.Version] = i
		}
		// Mark removed versions.
		for i := range tmpl.Status.Versions {
			if _, ok := tmplVersReverseIndexer[tmpl.Status.Versions[i].Version]; ok {
				delete(tmplVersReverseIndexer, tmpl.Status.Versions[i].Version)
				continue
			}
			tmpl.Status.Versions[i].Removed = true
		}
		// Append new versions.
		newTmplVersIndexes := maps.Values(tmplVersReverseIndexer)
		sort.Ints(newTmplVersIndexes)
		for _, i := range newTmplVersIndexes {
			tmpl.Status.Versions = append(tmpl.Status.Versions, tmplVers[i])
		}
	}

	return tmpl, nil
}

// genTemplateVersions retrieves template versions from a git repository.
func genTemplateVersions(
	ctx context.Context,
	obj *walruscore.Template,
	versions []string,
	versionSchema map[string]*api.SchemaGroup,
) ([]walruscore.TemplateVersion, error) {
	if len(versions) == 0 {
		return nil, nil
	}

	var (
		logger = klog.NewStandardLogger("WARNING")
		tvs    = make([]walruscore.TemplateVersion, 0, len(versionSchema))
	)

	su, err := sourceurl.ParseURLToSourceURL(obj.Spec.VCSRepository.URL)
	if err != nil {
		return nil, err
	}

	for i := range versions {
		version := versions[i]
		schema, ok := versionSchema[version]
		if !ok {
			logger.Printf("%s/%s version: %s version schema not found", obj.Namespace, obj.Name, version)
			continue
		}

		var u string
		{
			link, err := url.Parse(su.Link)
			if err != nil {
				return nil, err
			}
			link.RawQuery = url.Values{"ref": []string{version}}.Encode()
			u = link.String()
		}

		// Generate template version.
		tv, err := kubehelper.GenTemplateVersion(ctx, u, version, obj, schema)
		if err != nil {
			return nil, err
		}

		tvs = append(tvs, *tv)
	}

	return tvs, nil
}
