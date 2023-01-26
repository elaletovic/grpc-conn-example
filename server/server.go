package server

import (
	"github.com/elaletovic/grpc-conn-example/communicator"
)

type grpcServer struct {
	communicator.UnimplementedCommunicatorServer
}
