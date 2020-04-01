package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoTunerOptions struct {
	l          hexa.Logger
	t          hexa.Translator
	uf         hecho.UserFinderBySub
	userSDK    hexa.UserSDK
	ctxCreator hecho.CtxCreator
}

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, cfg hexa.Config, options EchoTunerOptions) {

	e.HideBanner = true

	e.Logger = hecho.HexaToEchoLogger(cfg, options.l)

	e.Debug = cfg.GetBool("debug")
	// Set the error handler.
	e.HTTPErrorHandler = hecho.HTTPErrorHandler(options.l, options.t, e.Debug)

	var currentUserMiddleware echo.MiddlewareFunc
	if options.uf == nil {
		currentUserMiddleware = hecho.CurrentUserWithoutFetch(options.userSDK)
	} else {
		currentUserMiddleware = hecho.CurrentUser(options.uf, options.userSDK)
	}

	// CORS HEADERS
	e.Use(middleware.CORSWithConfig(hecho.CorsConfig(cfg)))

	// Logger each request
	e.Use(middleware.Logger())

	// Recover recover each panic and pass to the cho error handler
	e.Use(hecho.Recover())

	// RequestID set requestID on each request that has blank request id.
	e.Use(hecho.RequestID())

	// CorrelationID set X-Correlation-ID value.
	e.Use(hecho.CorrelationID())

	// Optional JWT checker : check if exists
	//header => verify, otherwise skip it.
	e.Use(hecho.JWT(hexa.Secret(cfg.GetString("SECRET"))))

	// Set user in each request context.
	e.Use(currentUserMiddleware)

	// HexaContext set hexa context on each request.
	e.Use(hecho.HexaContext(options.ctxCreator, options.l, options.t))

	// SetContextLogger set the echo logger on each echo's context.
	e.Use(hecho.SetContextLogger(cfg))
}
