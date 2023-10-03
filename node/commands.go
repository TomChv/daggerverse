package main

import "context"

// Run returns the container with the command set.
func (n *Node) Run(args []string) *Container {
	return n.Ctr.WithExec(args)
}

// Start returns the output of the code executed.
func (n *Node) Start(ctx context.Context) (string, error) {
	return n.Run([]string{"start"}).Stdout(ctx)
}

// Lint returns the output of the lint command.
func (n *Node) Lint(ctx context.Context) (string, error) {
	return n.Run([]string{"lint"}).Stdout(ctx)
}

// Test returns the result of the test executed.
func (n *Node) Test(ctx context.Context) (string, error) {
	return n.Run([]string{"test"}).Stdout(ctx)
}

// Build returns the directory of the source built.
func (n *Node) Build(ctx context.Context) (*Directory, error) {
	result, err := n.Run([]string{"build"}).Sync(ctx)
	if err != nil {
		return nil, err
	}

	return result.Directory("dist"), nil
}

type InstallOpts struct {
	Pkg []string `doc:"Package to additionally install"`
}

// Install adds given package.
func (n *Node) Install(opts InstallOpts) *Node {
	cmd := append([]string{"install"}, opts.Pkg...)

	return n.WithContainer(n.Run(cmd))
}
