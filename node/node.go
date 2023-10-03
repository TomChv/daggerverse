package main

import (
	"context"
	"fmt"
)

type Node struct {
	Version string
	Cache   bool
}

func (n *Node) WithVersion(version string) *Node {
	n.Version = version

	return n
}

func (n *Node) Container() *Container {
	version := "latest"
	if n.Version != "" {
		version = n.Version
	}

	ctr := dag.Container().
		From(fmt.Sprintf("node:%s", version)).
		WithEntrypoint([]string{"node"})

	return ctr
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
