package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/internal/logs"
	"github.com/hanshal101/snapwall/models"
	snapwall "github.com/hanshal101/snapwall/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	snapwall.UnimplementedSenderServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error in loading '.env': %v\n", err)
		return
	}
	clickhouse.InitClickhouse(context.Background())
}

func (s *Server) Send(
	stream snapwall.Sender_SendServer,
) error {
	for {
		inp, err := stream.Recv()
		if err != nil {
			return err
		}

		log.Printf("Storing in Clickhouse: %v\n", inp)
		if err := logs.StoreLogs(context.Background(), &models.Log{
			Time:        inp.Time.AsTime(),
			Source:      inp.Source,
			Destination: inp.Destination,
			Type:        inp.Type,
			Port:        inp.Port,
			Protocol:    inp.Protocol,
		}); err != nil {
			log.Printf("Error in storing logs:\n Log: %v\n Error: %v\n", inp, err)
			return err
		}

		resp := &snapwall.ServiceResponse{
			Time:        inp.Time,
			Source:      inp.Source,
			Destination: inp.Destination,
			Type:        inp.Type,
			Port:        inp.Port,
			Protocol:    inp.Protocol,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	snapwall.RegisterSenderServer(s, &Server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
