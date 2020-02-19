package ktuner

import (
	"github.com/Kamva/elogrus/v4"
	"github.com/Kamva/gutil"
	"github.com/Kamva/kitty"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, config kitty.Config) {
	logger := gutil.Must(tuneLogrus(config)).(*logrus.Logger)

	// Set echo logger
	e.Logger = elogrus.GetEchoLogger(logger)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
