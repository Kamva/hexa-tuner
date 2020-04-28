package huner

import (
	"github.com/Kamva/hexa"
	"github.com/Kamva/hexa-rpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// GRPCConn returns new instance of the gRPC connection with your config to use in client
func GRPCConn(serverAddr string, cei hexa.ContextExporterImporter) (*grpc.ClientConn, error) {
	hexaCtxInt := hrpc.NewHexaContextInterceptor(cei)
	unaryInt := grpc.WithUnaryInterceptor(hexaCtxInt.UnaryClientInterceptor)
	// TODO: Init metric API and distributed tracing here.

	return grpc.Dial(serverAddr, grpc.WithInsecure(), unaryInt)
}

// TuneGRPCServer returns new instance of the tuned gRPC Server to server requests to services
func TuneGRPCServer(cei hexa.ContextExporterImporter, cfg hexa.Config, l hexa.Logger) (*grpc.Server, error) {
	loggerOptions := hrpc.DefaultLoggerOptions(cfg.GetBool("DEBUG"))

	// Replace gRPC logger with hexa logger
	grpclog.SetLoggerV2(hrpc.NewLogger(l, cfg))

	intChain := grpc_middleware.ChainUnaryServer(
		// Set Hexa context interceptor
		hrpc.NewHexaContextInterceptor(cei).UnaryServerInterceptor,
		// Set request logger
		hrpc.NewRequestLogger(l).UnaryServerInterceptor(loggerOptions),
	)
	return grpc.NewServer(grpc.UnaryInterceptor(intChain)), nil
}
