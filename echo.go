package ktuner

import (
	"github.com/Kamva/kitty"
	kecho "github.com/Kamva/kitty-echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, config kitty.Config, logger kitty.Logger, t kitty.Translator, uf kecho.UserFinderByJwtSub) {
	// Set echo logger
	e.Logger = kecho.KittyLoggerToEchoLogger(logger)

	// Set the error handler.
	e.HTTPErrorHandler = kecho.HTTPErrorHandler

	// Logger each request
	e.Use(middleware.Logger())

	// Recover recover each panic and pass to the cho error handler
	e.Use(middleware.Recover())

	// RequestID set requestID on each request that has blank request id.
	e.Use(middleware.RequestID())

	// Optional JWT checker : check if exists
	//header => verify, otherwise skip it.
	e.Use(kecho.JWT(config.GetString("SECRET")))

	// Set user in each request context.
	e.Use(kecho.CurrentUser(uf))

	// KittyContext set kitty context on each request.
	e.Use(kecho.KittyContext(logger, t))

}
