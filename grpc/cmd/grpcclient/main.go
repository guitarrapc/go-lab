package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	pb "github.com/RichardLaos/grpc-stream/notify"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func main() {
	keepaliveParams := keepalive.ClientParameters{
		Time:                5 * time.Second,
		Timeout:             5 * time.Second,
		PermitWithoutStream: true,
	}

	host := os.Getenv("SERVER_ENDPOINT")
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		host = "localhost:50051"
	}

	credential := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
	if !strings.HasSuffix(host, ":443") {
		credential = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	conn, err := grpc.Dial(host,
		credential,
		grpc.WithKeepaliveParams(keepaliveParams))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := client{
		client: pb.NewNotifyClient(conn),
		id:     "notify-client-1",
	}

	c.listen()
}

type client struct {
	client pb.NotifyClient
	id     string
}

func (c *client) listen() {
	var err error
	var stream pb.Notify_SubscribeClient

	for {
		if stream == nil {
			if stream, err = c.client.Subscribe(context.Background(), &pb.Request{DeviceId: c.id}); err != nil {
				log.Printf("failed to subscribe: %v", err)
				<-time.After(5 * time.Second)
				// Retry on failure
				continue
			}
		}
		response, err := stream.Recv()
		if err != nil {
			log.Printf("failed to receive: %v", err)
			// Clearing the stream will force the client to resubscribe on next iteration
			stream = nil
			<-time.After(5 * time.Second)
			// Retry on failure
			continue
		}
		log.Printf("device %s got response: %q", c.id, response.Message)
	}
}
