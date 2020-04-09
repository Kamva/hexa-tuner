package huner

import (
	"github.com/Kamva/hexa"
	hgrpc "github.com/Kamva/hexa-grpc"
	"google.golang.org/grpc"
)

// GRPCConn returns new instance of the gRPC connection with your config to use in client
func GRPCConn(serverAddr string, cei hexa.ContextExporterImporter) (*grpc.ClientConn, error) {
	hexaCtxtInt := hgrpc.NewHexaContextInterceptor(cei)
	unaryInt := grpc.WithUnaryInterceptor(hexaCtxtInt.UnaryClientInterceptor)

	return grpc.Dial(serverAddr, grpc.WithInsecure(), unaryInt)
}

// GRPCServer returns new instance of the gRPC Server to server requests to services
func GRPCServer(cei hexa.ContextExporterImporter) (*grpc.Server, error) {
	hexaCtxInt := hgrpc.NewHexaContextInterceptor(cei)
	return grpc.NewServer(grpc.UnaryInterceptor(hexaCtxInt.UnaryServerInterceptor)), nil
}
