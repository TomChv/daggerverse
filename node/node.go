package main

import (
	"context"
	"fmt"
)

type Node struct {
	Version string
	Cache   bool
}

func (n *Node) WithVersion(_ context.Context, version string) (*Node, error) {
	n.Version = version

	return n, nil
}

func (n *Node) Container(_ context.Context) (*Container, error) {
	version := "latest"
	if n.Version != "" {
		version = n.Version
	}

	ctr := dag.Container().
		From(fmt.Sprintf("node:%s", version)).
		WithEntrypoint([]string{"node"})

	return ctr, nil
}

// NodeCache adds node_modules into dagger cache.
// This expects that container workdir is set to source directory.
func (ctr *Container) NodeCache(ctx context.Context) (*Container, error) {
	workdir, err := ctr.Workdir(ctx)
	if err != nil {
		return nil, err
	}

	return ctr.
		WithMountedCache(
			fmt.Sprintf("%s/workdir", workdir),
			dag.CacheVolume("node-modules")), nil
}
