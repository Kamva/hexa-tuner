package ktuner

import (
	"github.com/Kamva/kitty"
	"github.com/Kamva/kitty-job"
	 "github.com/Kamva/kitty-job/faktory"
	"github.com/Kamva/tracer"
	faktory "github.com/contribsys/faktory/client"
	faktoryworker "github.com/contribsys/faktory_worker_go"
)

// NewFaktoryJobsDriver generate new faktory driver for kitty jobs.
func NewFaktoryJobsDriver(poolSize int) (kjob.Jobs, error) {
	p, err := faktory.NewPool(poolSize)

	if err != nil {
		return nil, tracer.Trace(err)
	}

	return kfaktory.NewFaktoryJobsDriver(p), nil
}

// NewFaktoryWorkerDriver generate new faktory driver for the kitty worker.
func NewFaktoryWorkersDriver(uf kitty.UserFinder, concurrency int, l kitty.Logger, t kitty.Translator) (kjob.Worker, error) {
	mgr := faktoryworker.NewManager()
	worker := kfaktory.NewFaktoryWorkerDriver(mgr, uf, l, t)
	err := worker.Concurrency(concurrency)

	return worker, tracer.Trace(err)
}
