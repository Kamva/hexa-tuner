package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa/hexalogger"
	"github.com/Kamva/logrus-kit"
	_ "github.com/Kamva/logrus-kit/logrusbase"
	"github.com/Kamva/tracer"
	"github.com/sirupsen/logrus"
)

// NewLogger return new instance of hexa logger service.
func NewLogger(config hexa.Config) (hexa.Logger, error) {
	logger, err := tuneLogrus(config)

	if err != nil {
		return nil, tracer.Trace(err)
	}

	return hexalogger.NewLogrusDriver(logger.WithFields(nil)), nil
}

func tuneLogrus(config hexa.Config) (*logrus.Logger, error) {
	logger := logrus.New()

	err := logruskit.Tune(logger, logruskit.Config(config))

	return logger, tracer.Trace(err)
}
