package main

import (
	"log"
	"net"

	s "github.com/elaletovic/grpc-conn-example/server/internal"
)

func main() {
	log.Println("Starting listening on port 6005")
	port := ":6005"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)
	srv := s.NewServer()

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
