package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/RichardLaos/grpc-stream/notify"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	pb.UnimplementedNotifyServer
	subscribers sync.Map
}

type subscriber struct {
	stream pb.Notify_SubscribeServer
	done   chan<- bool
}

func (s *server) Subscribe(request *pb.Request, stream pb.Notify_SubscribeServer) error {
	log.Printf("client subscribed: %s", request.DeviceId)

	done := make(chan bool)
	s.subscribers.Store(request.DeviceId, subscriber{stream: stream, done: done})

	ctx := stream.Context()
	for {
		select {
		case <-done:
			return nil
		case <-ctx.Done():
			// client disconnected
			return nil
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    10 * time.Second,
			Timeout: 15 * time.Second,
		}),
	)

	ns := &server{}
	go ns.sendNotifications()

	pb.RegisterNotifyServer(s, ns)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) sendNotifications() {
	for {
		time.Sleep(10 * time.Minute)

		var unsubscribe []string

		s.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(string)
			if !ok {
				return false
			}

			sub, ok := v.(subscriber)
			if !ok {
				return false
			}
			// Send data over the gRPC stream to the client
			if err := sub.stream.Send(&pb.Notification{Message: fmt.Sprintf("Notification for: %s", id)}); err != nil {
				select {
				case sub.done <- true:
				default:
					// avoid blocking
				}
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})

		// Unsubscribe
		for _, id := range unsubscribe {
			s.subscribers.Delete(id)
		}
	}
}
