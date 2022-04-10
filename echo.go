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
}

type EchoConfigs struct {
	MetricsConfig            hecho.MetricsConfig
	TracingConfig            hecho.TracingConfig
	JWTConfig                middleware.JWTConfig
	JwtClaimAuthorizerConfig hecho.JwtClaimAuthorizerConfig
	Debug                    bool
	EchoLogLevel             string
	CORS                     middleware.CORSConfig
}

// TuneEcho tune echo framework.
func TuneEcho(e *echo.Echo, cfg EchoConfigs, o EchoTunerOptions) {

	e.HideBanner = true
	e.Logger = hecho.HexaToEchoLogger(o.Logger, cfg.EchoLogLevel)
	e.Debug = cfg.Debug
	e.HTTPErrorHandler = hecho.HTTPErrorHandler(o.Logger, o.Translator, e.Debug)

	currentUserMiddleware := hecho.CurrentUser(o.UserFinder)
	if o.UserFinder == nil {
		currentUserMiddleware = hecho.CurrentUserWithoutFetch()
	}

	// CORS HEADERS
	e.Use(middleware.CORSWithConfig(cfg.CORS))

	// Log each request
	e.Use(middleware.Logger())

	e.Use(hecho.Metrics(cfg.MetricsConfig))
	e.Use(hecho.Tracing(cfg.TracingConfig))

	// Recover recovers each panic and returns its to the echo error handler
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
	e.Use(hecho.HexaContext(o.Logger, o.Translator))

	// SetContextLogger set the echo logger on each echo's context.
	e.Use(hecho.SetContextLogger(cfg.EchoLogLevel))

	// Add more data to each trace span:
	e.Use(hecho.TracingDataFromUserContext())
}
