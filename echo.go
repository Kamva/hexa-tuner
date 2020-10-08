package huner

import (
	"github.com/kamva/hexa"
	"github.com/kamva/hexa-echo"
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

type EchoConfigs struct {
	JWTConfig                middleware.JWTConfig
	JwtClaimAuthorizerConfig hecho.JwtClaimAuthorizerConfig
	Debug                    bool
	EchoLogLevel             string
	AllowOrigins             []string
	AllowHeaders             []string
	AllowMethods             []string
}

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, cfg EchoConfigs, o EchoTunerOptions) {

	e.HideBanner = true

	e.Logger = hecho.HexaToEchoLogger(o.Logger, cfg.EchoLogLevel)

	e.Debug = cfg.Debug
	// Set the error handler.
	e.HTTPErrorHandler = hecho.HTTPErrorHandler(o.Logger, o.Translator, e.Debug)

	var currentUserMiddleware echo.MiddlewareFunc
	if o.UserFinder == nil {
		currentUserMiddleware = hecho.CurrentUserWithoutFetch(o.UserSDK)
	} else {
		currentUserMiddleware = hecho.CurrentUser(o.UserFinder, o.UserSDK)
	}

	// CORS HEADERS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: cfg.AllowMethods,
		AllowHeaders: cfg.AllowHeaders,
	}))

	// Logger each request
	e.Use(middleware.Logger())

	// Recover recover each panic and pass to the cho error handler
	e.Use(hecho.Recover())

	// RequestID set requestID on each request that has blank request id.
	e.Use(hecho.RequestID())

	// CorrelationID set X-Correlation-ID value.
	e.Use(hecho.CorrelationID())

	// Optional JWT checker checks if exists "Authorization" header, so verify it, otherwise skip.
	e.Use(middleware.JWTWithConfig(cfg.JWTConfig))
	// JWT authorizer to authorize jwt claim
	e.Use(hecho.JwtClaimAuthorizer(cfg.JwtClaimAuthorizerConfig))


	// Set user in each request context.
	e.Use(currentUserMiddleware)

	// HexaContext set hexa context on each request.
	e.Use(hecho.HexaContext(o.CtxCreator, o.Logger, o.Translator))

	// SetContextLogger set the echo logger on each echo's context.
	e.Use(hecho.SetContextLogger(cfg.EchoLogLevel))
}
