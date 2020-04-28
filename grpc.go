package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-rpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type GRPCServerTunerOptions struct {
	cei hexa.ContextExporterImporter
	cfg hexa.Config
	l   hexa.Logger
	t   hexa.Translator
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
	loggerOptions := hrpc.DefaultLoggerOptions(o.cfg.GetBool("DEBUG"))

	// Replace gRPC logger with hexa logger
	grpclog.SetLoggerV2(hrpc.NewLogger(o.l, o.cfg))

	intChain := grpc_middleware.ChainUnaryServer(
		// Hexa context interceptor
		hrpc.NewHexaContextInterceptor(o.cei).UnaryServerInterceptor,

		// Request logger
		hrpc.NewRequestLogger(o.l).UnaryServerInterceptor(loggerOptions),

		// Hexa error interceptor (Must be last interceptor)
		hrpc.NewErrorInterceptor().UnaryServerInterceptor(o.t),
	)
	return grpc.NewServer(grpc.UnaryInterceptor(intChain)), nil
}
