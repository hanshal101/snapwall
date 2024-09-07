package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/database/psql"
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
	psql.InitDB()
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

		fmt.Println("matching policy..........")
		if matchPolicy(inp) {
			inp.Severity = "HIGH"
		} else {
			inp.Severity = "LOW"
		}

		fmt.Println("SEV:", inp.Severity)

		log.Printf("Storing in Clickhouse: %v\n", inp)
		if err := logs.StoreLogs(context.Background(), &models.Log{
			Time:        inp.Time.AsTime(),
			Source:      inp.Source,
			Destination: inp.Destination,
			Type:        inp.Type,
			Port:        inp.Port,
			Protocol:    inp.Protocol,
			Severity:    inp.Severity,
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
			Severity:    inp.Severity,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func matchPolicy(inp *snapwall.ServiceRequest) bool {
	var policies []models.Policy
	if err := psql.DB.Preload("IPs").Preload("Ports").Find(&policies).Error; err != nil {
		log.Printf("Error in fetching policies: %v", err)
		return false
	}

	for _, policy := range policies {
		for _, ips := range policy.IPs {
			if ips.Address == inp.Source {
				for _, port := range policy.Ports {
					if port.Number == inp.Port {
						log.Println("INTRUDER FOUND !!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
						return true
					}
				}
			}
		}
	}

	return false
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
