package platformk8s

import (
	"bufio"
	"context"
	"fmt"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/utils/pointer"

	"github.com/seal-io/seal/pkg/platform/operator"
	"github.com/seal-io/seal/pkg/platformk8s/key"
	"github.com/seal-io/seal/pkg/platformk8s/kube"
)

// Log implements operator.Operator.
func (op Operator) Log(ctx context.Context, k string, opts operator.LogOptions) error {
	// parse key.
	ns, pn, ct, cn, ok := key.Decode(k)
	if !ok {
		return fmt.Errorf("failed to parse given key: %q", k)
	}

	// confirm.
	var cli, err = coreclient.NewForConfig(op.RestConfig)
	if err != nil {
		return fmt.Errorf("error creating kubernetes client: %w", err)
	}
	p, err := cli.Pods(ns).
		Get(ctx, pn, meta.GetOptions{ResourceVersion: "0"}) // non quorum read
	if err != nil {
		return fmt.Errorf("error getting kubernetes pod %s/%s: %w", ns, pn, err)
	}
	if !kube.IsContainerExisted(p, kube.Container{Type: ct, Name: cn}) {
		return fmt.Errorf("given %s container %s is not ownered by %s/%s pod", ct, cn, ns, pn)
	}

	// stream.
	var stmOpts = &core.PodLogOptions{
		Container:    cn,
		Follow:       kube.IsContainerRunning(p, kube.Container{Type: ct, Name: cn}),
		Previous:     opts.Previous,
		SinceSeconds: opts.SinceSeconds,
		Timestamps:   opts.Timestamps,
	}
	if opts.Tail {
		stmOpts.TailLines = pointer.Int64(10)
	}
	stm, err := cli.Pods(ns).
		GetLogs(pn, stmOpts).
		Stream(ctx)
	if err != nil {
		return fmt.Errorf("failed to create log stream: %w", err)
	}
	defer func() { _ = stm.Close() }()
	var r = bufio.NewReader(stm)
	var w = opts.Out
	for {
		var bs []byte
		bs, err = r.ReadBytes('\n')
		if err != nil {
			if isTrivialError(err) {
				err = nil
			}
			break
		}
		_, err = w.Write(bs)
		if err != nil {
			if isTrivialError(err) {
				err = nil
			}
			break
		}
	}
	if err != nil {
		return fmt.Errorf("error streaming log: %w", err)
	}
	return nil
}