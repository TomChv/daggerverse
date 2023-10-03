package main

import "fmt"

type Deno struct {
	Ctr *Container
}

// WithVersion returns Deno container with given image version.
func (d *Deno) WithVersion(version string) *Deno {
	if d.Ctr == nil {
		d.Ctr = dag.Container()
	}

	d.Ctr = d.Ctr.
		From(fmt.Sprintf("denoland/deno:%s", version)).
		WithEntrypoint([]string{"deno"}).
		WithMountedCache("/root/.cache/deno", dag.CacheVolume("deno-cache"))

	return d
}

// WithContainer returns Deno container with the given container.
func (d *Deno) WithContainer(ctr *Container) *Deno {
	d.Ctr = ctr

	return d
}

// Container returns Deno container.
func (d *Deno) Container() *Container {
	return d.Ctr
}

func (d *Deno) WithSource(directory *Directory) *Deno {
	workdir := "/src"

	d.Ctr = d.Ctr.
		WithWorkdir(workdir).
		WithMountedDirectory(workdir, directory)

	return d
}
