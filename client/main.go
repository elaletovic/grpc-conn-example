package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	comm "github.com/elaletovic/grpc-conn-example/communicator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:6005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := comm.NewCommunicatorClient(conn)

	go checkConnection(conn)

	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			msg := &comm.MessageRequest{Id: fmt.Sprintf("%d", rand.Int63()), Payload: "Message"}

			response, err := client.Send(context.Background(), msg)
			if err != nil {
				log.Println("server error", err)
			} else {
				log.Println("response", response)
			}
		}
	}
}

func checkConnection(conn *grpc.ClientConn) {
	for {
		time.Sleep(2 * time.Second)

		log.Println("connection state is", conn.GetState())
	}
}
