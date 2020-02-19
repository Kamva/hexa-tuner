package ktuner

import (
	"github.com/Kamva/kitty"
	"github.com/Kamva/kitty/kittylogger"
	"github.com/Kamva/logrus-kit"
	_ "github.com/Kamva/logrus-kit/logrusbase"
	"github.com/sirupsen/logrus"
)

// NewLogger return new instance of kitty logger service.
func NewLogger(config kitty.Config) (kitty.Logger, error) {
	logger, err := tuneLogrus(config)

	if err != nil {
		return nil, err
	}

	return kittylogger.NewLogrusDriver(logger.WithFields(nil)), nil
}

func tuneLogrus(config kitty.Config) (*logrus.Logger, error) {
	logger := logrus.New()

	err := logruskit.Tune(logger, logruskit.Config(config))

	return logger, err
}
