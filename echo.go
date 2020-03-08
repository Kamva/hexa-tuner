package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, config hexa.Config, l hexa.Logger, t hexa.Translator, uf hecho.UserFinderByJwtSub) {

	e.HideBanner = true

	e.Logger = hecho.HexaToEchoLogger(l)

	e.Debug = config.GetBool("debug")
	// Set the error handler.
	e.HTTPErrorHandler = hecho.HTTPErrorHandler(l, t, e.Debug)

	// Logger each request
	e.Use(middleware.Logger())

	// Recover recover each panic and pass to the cho error handler
	e.Use(middleware.Recover())

	// RequestID set requestID on each request that has blank request id.
	e.Use(hecho.RequestID())

	// CorrelationID set X-Correlation-ID value.
	e.Use(hecho.CorrelationID())

	// Optional JWT checker : check if exists
	//header => verify, otherwise skip it.
	e.Use(hecho.JWT(hexa.Secret(config.GetString("SECRET"))))

	// Set user in each request context.
	e.Use(hecho.CurrentUser(uf))

	// HexaContext set hexa context on each request.
	e.Use(hecho.HexaContext(l, t))

}
