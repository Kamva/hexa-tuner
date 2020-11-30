package huner

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kamva/gutil"
	"github.com/kamva/hexa"
	"github.com/kamva/hexa-rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// GRPCServerTunerOptions contains options needed to tune a gRPC server
type GRPCServerTunerOptions struct {
	ContextEI  hexa.ContextExporterImporter
	Logger     hexa.Logger
	Translator hexa.Translator
}

type GRPCConfigs struct {
	Secret       string
	Debug        bool
	LogVerbosity int `json:"log_verbosity" yaml:"log_verbosity"`
}

// Must returns new instance of the gRPC connection with your config to use in client
// it will panic any error.
func MustGRPCConn(serverAddr string, cei hexa.ContextExporterImporter) *grpc.ClientConn {
	return gutil.Must(GRPCConn(serverAddr, cei)).(*grpc.ClientConn)
}

// GRPCConn returns new instance of the gRPC connection with your config to use in client
func GRPCConn(serverAddr string, cei hexa.ContextExporterImporter) (*grpc.ClientConn, error) {
	unaryInt := grpc.WithChainUnaryInterceptor(
		// Hexa error interceptor (convert gRPC status to hexa error)
		hrpc.NewErrorInterceptor().UnaryClientInterceptor(),
		// Hexa context interceptor
		hrpc.NewHexaContextInterceptor(cei).UnaryClientInterceptor,
	)
	// TODO: Init metric API and distributed tracing here.

	return grpc.Dial(serverAddr, grpc.WithInsecure(), unaryInt)
}

// TuneGRPCServer returns new instance of the tuned gRPC Server to server requests to services
func TuneGRPCServer(cfg GRPCConfigs, o GRPCServerTunerOptions) (*grpc.Server, error) {
	loggerOptions := hrpc.DefaultLoggerOptions(cfg.Debug)

	errOptions := hrpc.ErrInterceptorOptions{
		Logger:       o.Logger,
		Translator:   o.Translator,
		ReportErrors: true,
	}

	// Replace gRPC logger with hexa logger
	grpclog.SetLoggerV2(hrpc.NewLogger(o.Logger, cfg.LogVerbosity))

	intChain := grpc_middleware.ChainUnaryServer(
		// Hexa context interceptor
		hrpc.NewHexaContextInterceptor(o.ContextEI).UnaryServerInterceptor,
		// Request logger
		hrpc.NewRequestLogger(o.Logger).UnaryServerInterceptor(loggerOptions),
		// Hexa error interceptor
		hrpc.NewErrorInterceptor().UnaryServerInterceptor(errOptions),
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(hrpc.RecoverHandler)),
	)

	return grpc.NewServer(grpc.UnaryInterceptor(intChain)), nil
}
