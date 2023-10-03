package main

import (
	"context"
	"fmt"
)

type Node struct{}

// Base stands as the basic node container image.
func (n *Node) Base(ctx context.Context, version string) (*Container, error) {
	address := fmt.Sprintf("node:%s", version)

	return dag.
		Container().
		From(address).
		WithEntrypoint([]string{"node"}).
		Sync(ctx)
}

// Version returns this container version
func (n *Node) Version(_ context.Context, ctr *Container) (*Container, error) {
	return ctr, nil
}

// Version returns the function of the image
func (ctr *Container) Version(ctx context.Context) (string, error) {
	return ctr.WithExec([]string{"-v"}).Stdout(ctx)
}
