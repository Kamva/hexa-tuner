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
	errNilConfig        = errors.New("config is Nil in the service container")
	errNilLogger        = errors.New("logger is Nil in the service container")
	errNilTranslator    = errors.New("translator is Nil in the service container")
	errNilHealthChecker = errors.New("healthReporter is nil in the service container")
	errNilJobs          = errors.New("jobs is nil in the service container")
	errNilEmitter       = errors.New("emitter is nil in the service container")
	errNilArranger      = errors.New("arranger is nil in the service container")
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

		Config() hexa.Config
		Logger() hexa.Logger
		Translator() hexa.Translator
		HealthReporter() hexa.HealthReporter
		Jobs() hjob.Jobs
		Emitter() hevent.Emitter
		Arranger() arranger.Arranger
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
	}
)

// SetConfig sets the config service.
func (p *baseServiceContainer) SetConfig(config hexa.Config) {
	p.config = config
}

// SetLogger sets the logger service.
func (p *baseServiceContainer) SetLogger(logger hexa.Logger) {
	p.log = logger
}

// SetTranslator sets the translator service.
func (p *baseServiceContainer) SetTranslator(translator hexa.Translator) {
	p.translator = translator
}

// SetHealthReporter sets the healthReporter service.
func (p *baseServiceContainer) SetHealthReporter(reporter hexa.HealthReporter) {
	p.healthReporter = reporter
}

// SetJobs sets the Jobs service.
func (p *baseServiceContainer) SetJobs(jobs hjob.Jobs) {
	p.jobs = jobs
}

// SetEmitter sets the event emitter service.
func (p *baseServiceContainer) SetEmitter(emitter hevent.Emitter) {
	p.emitter = emitter
}

// SetArranger sets the arranger service.
func (p *baseServiceContainer) SetArranger(arranger arranger.Arranger) {
	p.arranger = arranger
}

// Config returns the config service.
func (p *baseServiceContainer) Config() hexa.Config {
	if p.must {
		gutil.PanicNil(p.config, errNilConfig)
	}

	return p.config
}

// Logger returns the logger service.
func (p *baseServiceContainer) Logger() hexa.Logger {
	if p.must {
		gutil.PanicNil(p.log, errNilLogger)
	}

	return p.log
}

// Translator returns the translator service.
func (p *baseServiceContainer) Translator() hexa.Translator {
	if p.must {
		gutil.PanicNil(p.translator, errNilTranslator)
	}
	return p.translator
}

// HealthReporter returns the healthReporter service.
func (p *baseServiceContainer) HealthReporter() hexa.HealthReporter {
	if p.must {
		gutil.PanicNil(p.healthReporter, errNilTranslator)
	}
	return p.healthReporter
}

// Jobs returns the jobs service.
func (p *baseServiceContainer) Jobs() hjob.Jobs {
	if p.must {
		gutil.PanicNil(p.jobs, errNilJobs)
	}

	return p.jobs
}

// Emitter returns the gate service.
func (p *baseServiceContainer) Emitter() hevent.Emitter {
	if p.must {
		gutil.PanicNil(p.emitter, errNilEmitter)
	}

	return p.emitter
}

// Arranger returns the gate service.
func (p *baseServiceContainer) Arranger() arranger.Arranger {
	if p.must {
		gutil.PanicNil(p.arranger, errNilArranger)
	}

	return p.arranger
}

// NewBaseServiceContainer returns new instance of the BaseServiceContainer.
func NewBaseServiceContainer(must bool) BaseServiceContainer {
	return &baseServiceContainer{
		must: must,
	}
}
