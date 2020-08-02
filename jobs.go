package huner

import (
	"github.com/kamva/hexa"
	"github.com/kamva/hexa-job"
	"github.com/kamva/hexa-job/faktory"
	"github.com/kamva/tracer"
	faktory "github.com/contribsys/faktory/client"
	faktoryworker "github.com/contribsys/faktory_worker_go"
)

// NewFaktoryJobsDriver generate new faktory driver for hexa jobs.
func NewFaktoryJobsDriver(ctxExporterImporter hexa.ContextExporterImporter, poolSize int) (hjob.Jobs, error) {
	p, err := faktory.NewPool(poolSize)

	if err != nil {
		return nil, tracer.Trace(err)
	}

	return hexafaktory.NewFaktoryJobsDriver(p, ctxExporterImporter), nil
}

// NewFaktoryWorkerDriver generate new faktory driver for the hexa worker.
func NewFaktoryWorkersDriver(ctxExporterImporter hexa.ContextExporterImporter, concurrency int) (hjob.Worker, error) {
	mgr := faktoryworker.NewManager()
	worker := hexafaktory.NewFaktoryWorkerDriver(mgr, ctxExporterImporter)
	err := worker.Concurrency(concurrency)

	return worker, tracer.Trace(err)
}
