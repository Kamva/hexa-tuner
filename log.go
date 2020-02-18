package ktuner

import (
	"github.com/Kamva/kitty"
	"github.com/Kamva/logrus-kit"
	_ "github.com/Kamva/logrus-kit/logrusbase"
	"github.com/sirupsen/logrus"
)

// NewLogger return new instance of logger.
func NewLogger(config kitty.Config) (*logrus.Logger, error) {
	logger := logrus.New()

	if err := logruskit.Tune(logger, logruskit.Config(config)); err != nil {
		return nil, err
	}

	return logger, nil
}
