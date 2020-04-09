package huner

import (
	"errors"
	"github.com/Kamva/gutil"
	"github.com/Kamva/hexa"
	hevent "github.com/Kamva/hexa-event"
	"github.com/Kamva/hexa-job"
)

var (
	errNilConfig     = errors.New("config is Nil in the pack")
	errNilLogger     = errors.New("logger is Nil in the pack")
	errNilTranslator = errors.New("translator is Nil in the pack")
	errNilJobs       = errors.New("jobs is nil in the pack")
	errNilGate       = errors.New("gate is nil in the pack")
)

// BaseServiceContainer is the base service container to use in each microservice.
type BaseServiceContainer interface {
	SetConfig(config hexa.Config)
	SetLogger(logger hexa.Logger)
	SetTranslator(translator hexa.Translator)
	SetJobs(jobs hjob.Jobs)
	SetGate(gate hexa.Gate)
	SetEmitter(emitter hevent.Emitter)
	Config() hexa.Config
	Logger() hexa.Logger
	Translator() hexa.Translator
	Jobs() hjob.Jobs
	Gate() hexa.Gate
	Emitter() hevent.Emitter
}

// baseServiceContainer contains all of services in one place to manage our services.
type baseServiceContainer struct {
	// must specify that should panic when user request
	// to get a nil service or just return nil value.
	must bool

	config     hexa.Config
	log        hexa.Logger
	translator hexa.Translator
	jobs       hjob.Jobs
	gate       hexa.Gate
	emitter    hevent.Emitter
}

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

// SetJobs sets the Jobs service.
func (p *baseServiceContainer) SetJobs(jobs hjob.Jobs) {
	p.jobs = jobs
}

// SetGate sets the Gate service.
func (p *baseServiceContainer) SetGate(gate hexa.Gate) {
	p.gate = gate
}

// SetGate sets the event emitter service.
func (p *baseServiceContainer) SetEmitter(emitter hevent.Emitter) {
	p.emitter = emitter
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

// Jobs returns the jobs service.
func (p *baseServiceContainer) Jobs() hjob.Jobs {
	if p.must {
		gutil.PanicNil(p.jobs, errNilJobs)
	}

	return p.jobs
}

// Gate returns the gate service.
func (p *baseServiceContainer) Gate() hexa.Gate {
	if p.must {
		gutil.PanicNil(p.gate, errNilGate)
	}

	return p.gate
}

// Emitter returns the gate service.
func (p *baseServiceContainer) Emitter() hevent.Emitter {
	if p.must {
		gutil.PanicNil(p.emitter, errNilGate)
	}

	return p.emitter
}

// NewBaseServiceContainer returns new instance of the BaseServiceContainer.
func NewBaseServiceContainer(must bool) BaseServiceContainer {
	return &baseServiceContainer{
		must: must,
	}
}
