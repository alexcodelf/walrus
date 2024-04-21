package vcssource

import (
	"context"
	"regexp"

	"github.com/drone/go-scm/scm"
	"github.com/seal-io/utils/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
	"github.com/seal-io/walrus/pkg/vcs"
	"github.com/seal-io/walrus/pkg/vcs/options"
)

func New() *Lister {
	return &Lister{}
}

// Lister is a vcs source list implement.
type Lister struct{}

// List returns a list of templates from the given catalog.
func (l *Lister) List(ctx context.Context, cat *walruscore.Catalog) ([]walruscore.Template, error) {
	logger := klog.Background().WithName("vcs").WithName("lister")

	var (
		source        = cat.Spec.VCSSource
		repos         []*scm.Repository
		filteredRepos []*scm.Repository
		tmpls         []walruscore.Template
	)

	// List.
	{
		opts, err := l.listOptions(ctx)
		if err != nil {
			return nil, err
		}

		repos, err = vcs.GetOrgRepos(ctx, source.Platform, source.URL, opts...)
		if err != nil {
			return nil, err
		}
		logger.Infof("found %d repositories in %s/%s before filtered", len(repos), cat.Namespace, cat.Name)
	}

	// Filtering.
	{
		var (
			includeReg *regexp.Regexp
			excludeReg *regexp.Regexp
			err        error
		)

		if filters := cat.Spec.Filters; filters != nil {
			if filters.IncludeExpression != "" {
				includeReg, err = regexp.Compile(filters.IncludeExpression)
				if err != nil {
					return nil, err
				}
			}
			if filters.ExcludeExpression != "" {
				excludeReg, err = regexp.Compile(filters.ExcludeExpression)
				if err != nil {
					return nil, err
				}
			}
		}

		for i := range repos {
			repo := repos[i]

			switch {
			case repo.Name == "":
				logger.Info("repository name is empty, skip")
				continue
			case len(validation.IsDNS1123Subdomain(repo.Name)) != 0:
				logger.Info("repository name is not a lowercase RFC 1123 subdomain name, skip", "repo", repo.Name)
				continue
			}

			if includeReg != nil && !includeReg.MatchString(repo.Name) {
				continue
			}
			if excludeReg != nil && excludeReg.MatchString(repo.Name) {
				continue
			}

			filteredRepos = append(filteredRepos, repo)
		}
	}

	// Generate Templates.
	{
		tmpls = make([]walruscore.Template, len(filteredRepos))
		for i := range filteredRepos {
			repo := filteredRepos[i]

			tmpl := walruscore.Template{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: cat.Namespace,
					Name:      kubehelper.NormalizeTemplateName(cat.Name, repo.Name),
					Labels:    GenWalrusBuiltinLabels(repo.Topics, cat.Name),
				},
				Spec: walruscore.TemplateSpec{
					TemplateFormat: cat.Spec.TemplateFormat,
					Description:    repo.Description,
					VCSRepository: &walruscore.VCSRepository{
						Platform: cat.Spec.VCSSource.Platform,
						URL:      repo.Link,
					},
				},
			}

			tmpls[i] = tmpl
		}

		logger.Infof("found %d repositories in %s/%s after filtered", len(tmpls), cat.Namespace, cat.Name)
	}

	return tmpls, nil
}

func (l *Lister) listOptions(ctx context.Context) ([]options.ClientOption, error) {
	opts := make([]options.ClientOption, 0)

	sid, err := systemsetting.ServeIdentify.Value(ctx)
	if err != nil {
		return nil, err
	}
	ua := version.GetUserAgent() + "; uuid=" + sid
	opts = append(opts, options.WithUserAgent(ua))

	tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
	if err != nil {
		return nil, err
	}

	if !tlsVerify {
		opts = append(opts, options.WithInsecureSkipVerify())
	}
	return opts, nil
}
