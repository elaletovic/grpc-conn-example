package server

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
	response := ""
	var err error
	if req != nil {
		log.Printf("message %s received, payload: %s\n", req.ID, req.Payload)
		response = fmt.Sprintf("message %s received successfully", req.ID)
	} else {
		response = "request cannot be nil"
		err = errors.New(response)
	}

	return response, err
}

func NewServer() *grpcServer {
	gsrv := grpc.NewServer()
	srv := grpcServer{}

	comm.RegisterCommunicatorServer(gsrv, &srv)

	return &gsrv
}
