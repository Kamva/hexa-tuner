package ktuner

import (
	"github.com/Kamva/elogrus/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func TuneEcho(e *echo.Echo, logger *logrus.Logger) {
	// Set echo logger
	e.Logger = elogrus.GetEchoLogger(logger)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
