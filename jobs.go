package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-job"
	"github.com/Kamva/hexa-job/faktory"
	"github.com/Kamva/tracer"
	faktory "github.com/contribsys/faktory/client"
	faktoryworker "github.com/contribsys/faktory_worker_go"
)

// NewFaktoryJobsDriver generate new faktory driver for hexa jobs.
func NewFaktoryJobsDriver(poolSize int) (hjob.Jobs, error) {
	p, err := faktory.NewPool(poolSize)

	if err != nil {
		return nil, tracer.Trace(err)
	}

	return hexafaktory.NewFaktoryJobsDriver(p), nil
}

// NewFaktoryWorkerDriver generate new faktory driver for the hexa worker.
func NewFaktoryWorkersDriver(uf hexa.UserFinder, concurrency int, l hexa.Logger, t hexa.Translator) (hjob.Worker, error) {
	mgr := faktoryworker.NewManager()
	worker := hexafaktory.NewFaktoryWorkerDriver(mgr, uf, l, t)
	err := worker.Concurrency(concurrency)

	return worker, tracer.Trace(err)
}
