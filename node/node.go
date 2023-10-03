package main

import (
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
// This path to target the node_modules directory
func (ctr *Container) NodeCache(path string) *Container {
	return ctr.
		WithMountedCache(
			fmt.Sprintf("%s/node_modules", path),
			dag.CacheVolume("node-modules"))
}
