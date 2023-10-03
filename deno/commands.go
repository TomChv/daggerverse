package main

import "context"

// Run returns the container with given command.
func (d *Deno) Run(args []string) *Container {
	return d.Ctr.WithExec(args)
}

// Task executes the given task and return its output.
func (d *Deno) Task(ctx context.Context, task string) (string, error) {
	return d.Run([]string{"task", task}).Stdout(ctx)
}

// Test returns the result of tests.
func (d *Deno) Test(ctx context.Context) (string, error) {
	return d.Run([]string{"test"}).Stdout(ctx)
}

// Lint returns the result of lint.
func (d *Deno) Lint(ctx context.Context) (string, error) {
	return d.Run([]string{"lint"}).Stdout(ctx)
}

// Fmt returns the deno container with fmt executed in it.
func (d *Deno) Fmt() *Deno {
	d.Ctr = d.Run([]string{"fmt"})

	return d
}

// Compile returns the container containing the compiled binary.
func (d *Deno) Compile(flags []string) *Deno {
	compileCmd := append([]string{"compile"}, flags...)

	d.Ctr = d.Run(compileCmd)

	return d
}
