package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoTunerOptions struct {
	Logger     hexa.Logger
	Translator hexa.Translator
	UserFinder hecho.UserFinderBySub
	UserSDK    hexa.UserSDK
	CtxCreator hecho.CtxCreator
}

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, cfg hexa.Config, options EchoTunerOptions) {

	e.HideBanner = true

	e.Logger = hecho.HexaToEchoLogger(cfg, options.Logger)

	e.Debug = cfg.GetBool("debug")
	// Set the error handler.
	e.HTTPErrorHandler = hecho.HTTPErrorHandler(options.Logger, options.Translator, e.Debug)

	var currentUserMiddleware echo.MiddlewareFunc
	if options.UserFinder == nil {
		currentUserMiddleware = hecho.CurrentUserWithoutFetch(options.UserSDK)
	} else {
		currentUserMiddleware = hecho.CurrentUser(options.UserFinder, options.UserSDK)
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
	// header, so verify it, otherwise skip.
	e.Use(hecho.JWT(hexa.Secret(cfg.GetString("SECRET"))))

	// Set user in each request context.
	e.Use(currentUserMiddleware)

	// HexaContext set hexa context on each request.
	e.Use(hecho.HexaContext(options.CtxCreator, options.Logger, options.Translator))

	// SetContextLogger set the echo logger on each echo's context.
	e.Use(hecho.SetContextLogger(cfg))
}
