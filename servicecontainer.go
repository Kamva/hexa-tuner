package huner

import (
	"errors"

	"github.com/kamva/gutil"
	"github.com/kamva/hexa"
	arranger "github.com/kamva/hexa-arranger"
	hevent "github.com/kamva/hexa-event"
	"github.com/kamva/hexa-job"
)

var (
	errNilConfig         = errors.New("config is Nil in the service container")
	errNilLogger         = errors.New("logger is Nil in the service container")
	errNilTranslator     = errors.New("translator is Nil in the service container")
	errNilHealthReporter = errors.New("healthReporter is nil in the service container")
	errNilJobs           = errors.New("jobs is nil in the service container")
	errNilEmitter        = errors.New("emitter is nil in the service container")
	errNilArranger       = errors.New("arranger is nil in the service container")
	errNilDLM            = errors.New("DLM is nil in the service container")
)

type (
	// BaseServiceContainer is the base service container to use in each microservice.
	BaseServiceContainer interface {
		SetConfig(config hexa.Config)
		SetLogger(logger hexa.Logger)
		SetTranslator(translator hexa.Translator)
		SetHealthReporter(reporter hexa.HealthReporter)
		SetJobs(jobs hjob.Jobs)
		SetEmitter(emitter hevent.Emitter)
		SetArranger(arranger arranger.Arranger)
		SetDLM(dlm hexa.DLM)

		Config() hexa.Config
		Logger() hexa.Logger
		Translator() hexa.Translator
		HealthReporter() hexa.HealthReporter
		Jobs() hjob.Jobs
		Emitter() hevent.Emitter
		Arranger() arranger.Arranger
		DLM() hexa.DLM
	}

	// baseServiceContainer contains all of services in one place to manage our services.
	baseServiceContainer struct {
		// must specify that should panic when user request
		// to get a nil service or just return nil value.
		must bool

		config         hexa.Config
		log            hexa.Logger
		translator     hexa.Translator
		healthReporter hexa.HealthReporter
		jobs           hjob.Jobs
		emitter        hevent.Emitter
		arranger       arranger.Arranger
		dlm            hexa.DLM
	}
)

// SetConfig sets the config service.
func (c *baseServiceContainer) SetConfig(config hexa.Config) {
	c.config = config
}

// SetLogger sets the logger service.
func (c *baseServiceContainer) SetLogger(logger hexa.Logger) {
	c.log = logger
}

// SetTranslator sets the translator service.
func (c *baseServiceContainer) SetTranslator(translator hexa.Translator) {
	c.translator = translator
}

// SetHealthReporter sets the healthReporter service.
func (c *baseServiceContainer) SetHealthReporter(reporter hexa.HealthReporter) {
	c.healthReporter = reporter
}

// SetJobs sets the Jobs service.
func (c *baseServiceContainer) SetJobs(jobs hjob.Jobs) {
	c.jobs = jobs
}

// SetEmitter sets the event emitter service.
func (c *baseServiceContainer) SetEmitter(emitter hevent.Emitter) {
	c.emitter = emitter
}

// SetArranger sets the arranger service.
func (c *baseServiceContainer) SetArranger(arranger arranger.Arranger) {
	c.arranger = arranger
}

// SetDLM sets the dlm(distributed lock manager) service.
func (c *baseServiceContainer) SetDLM(dlm hexa.DLM) {
	c.dlm = dlm
}

// Config returns the config service.
func (c *baseServiceContainer) Config() hexa.Config {
	if c.must {
		gutil.PanicNil(c.config, errNilConfig)
	}

	return c.config
}

// Logger returns the logger service.
func (c *baseServiceContainer) Logger() hexa.Logger {
	if c.must {
		gutil.PanicNil(c.log, errNilLogger)
	}

	return c.log
}

// Translator returns the translator service.
func (c *baseServiceContainer) Translator() hexa.Translator {
	if c.must {
		gutil.PanicNil(c.translator, errNilTranslator)
	}
	return c.translator
}

// HealthReporter returns the healthReporter service.
func (c *baseServiceContainer) HealthReporter() hexa.HealthReporter {
	if c.must {
		gutil.PanicNil(c.healthReporter, errNilHealthReporter)
	}
	return c.healthReporter
}

// Jobs returns the jobs service.
func (c *baseServiceContainer) Jobs() hjob.Jobs {
	if c.must {
		gutil.PanicNil(c.jobs, errNilJobs)
	}

	return c.jobs
}

// Emitter returns the gate service.
func (c *baseServiceContainer) Emitter() hevent.Emitter {
	if c.must {
		gutil.PanicNil(c.emitter, errNilEmitter)
	}

	return c.emitter
}

// Arranger returns the gate service.
func (c *baseServiceContainer) Arranger() arranger.Arranger {
	if c.must {
		gutil.PanicNil(c.arranger, errNilArranger)
	}

	return c.arranger
}

func (c *baseServiceContainer) DLM() hexa.DLM {
	return c.dlm
}

// NewBaseServiceContainer returns new instance of the BaseServiceContainer.
func NewBaseServiceContainer(must bool) BaseServiceContainer {
	return &baseServiceContainer{
		must: must,
	}
}
