package internal

import (
	"context"
	"errors"
	"fmt"
	"log"

	comm "github.com/elaletovic/grpc-conn-example/communicator"
	"google.golang.org/grpc"
)

type grpcServer struct {
	comm.UnimplementedCommunicatorServer
}

func (s *grpcServer) Send(ctx context.Context, req *comm.MessageRequest) (*comm.MessageResponse, error) {
	var response *comm.MessageResponse
	var err error
	if req != nil {
		log.Printf("message %s received, payload: %s\n", req.Id, req.Payload)
		response = &comm.MessageResponse{Status: fmt.Sprintf("message %s received successfully", req.Id)}
	} else {
		err = errors.New("request cannot be nil")
	}

	return response, err
}

func NewServer() *grpc.Server {
	gsrv := grpc.NewServer()
	srv := grpcServer{}

	comm.RegisterCommunicatorServer(gsrv, &srv)

	return gsrv
}
