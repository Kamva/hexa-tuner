package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-rpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// GRPCServerTunerOptions contains options needed to tune a gRPC server
type GRPCServerTunerOptions struct {
	ContextEI  hexa.ContextExporterImporter
	Config     hexa.Config
	Logger     hexa.Logger
	Translator hexa.Translator
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
func TuneGRPCServer(o GRPCServerTunerOptions) (*grpc.Server, error) {
	loggerOptions := hrpc.DefaultLoggerOptions(o.Config.GetBool("DEBUG"))

	errOptions := hrpc.ErrInterceptorOptions{
		Logger:       o.Logger,
		Translator:   o.Translator,
		ReportErrors: true,
	}

	// Replace gRPC logger with hexa logger
	grpclog.SetLoggerV2(hrpc.NewLogger(o.Logger, o.Config))

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
