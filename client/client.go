package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	snapwall "github.com/hanshal101/snapwall/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getBpfFilter(logType string) string {
	excludePort := "not port 5051 and not port 22 and not port 63262 and not port 62616"

	switch logType {
	case "http":
		return "tcp port 80 or tcp port 443 and " + excludePort
	case "tcp":
		return "tcp and " + excludePort
	case "udp":
		return "udp and " + excludePort
	case "all-scans":
		return "tcp[tcpflags] & (tcp-syn|tcp-fin|tcp-rst|tcp-ack) != 0 and " + excludePort
	case "icmp":
		return "icmp"
	default:
		return excludePort
	}
}

// func identifyService(port layers.TCPPort) string {
// 	switch port {
// 	case 80:
// 		return "HTTP"
// 	case 443:
// 		return "HTTPS"
// 	case 22:
// 		return "SSH"
// 	case 53:
// 		return "DNS"
// 	default:
// 		return "Unknown Service"
// 	}
// }

type Policy struct {
	Name  string
	Ports []string
	IPs   []string
	Type  bool
}

func main() {
	// policies := []Policy{
	// 	Policy{
	// 		Name:  "pol-1",
	// 		Ports: []string{":80", "443"},
	// 		IPs:   []string{"192.168.200.1"},
	// 		Type:  false,
	// 	},
	// }

	logType := flag.String("log", "all", "Type of logs to capture (http, tcp, udp, all-scans, icmp, all)")
	iface := flag.String("iface", "\\Device\\NPF_{A4770599-05C8-4E3F-8715-3D51E41B74BE}", "Your Network Packet destination")
	flag.Parse()

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := snapwall.NewSenderClient(conn)

	handle, err := pcap.OpenLive(*iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer handle.Close()

	filter := getBpfFilter(*logType)
	if filter != "" {
		err = handle.SetBPFFilter(filter)
		if err != nil {
			log.Fatalf("Failed to set BPF filter: %v", err)
		}
		fmt.Printf("BPF Filter Set: %s\n", filter)
	} else {
		fmt.Println("No specific filter applied. Capturing all traffic.")
	}

	localIPs, err := getLocalIPs()
	if err != nil {
		log.Fatalf("Failed to get local IPs: %v", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("Starting packet capture...")
	for packet := range packetSource.Packets() {
		networkLayer := packet.NetworkLayer()
		transportLayer := packet.TransportLayer()

		if networkLayer != nil && transportLayer != nil {
			srcIP, dstIP := networkLayer.NetworkFlow().Src().String(), networkLayer.NetworkFlow().Dst().String()

			direction := "Unknown"
			if contains(localIPs, srcIP) {
				direction = "Outgoing"
			} else if contains(localIPs, dstIP) {
				direction = "Incoming"
			}

			var port string
			var protocol string
			switch layer := transportLayer.(type) {
			case *layers.TCP:
				port = fmt.Sprintf("%d", layer.DstPort)
				protocol = "TCP"
			case *layers.UDP:
				port = fmt.Sprintf("%d", layer.DstPort)
				protocol = "UDP"
			case *layers.UDPLite:
				port = fmt.Sprintf("%d", layer.DstPort)
				protocol = "UDP"
			// case *layers.NortelDiscovery:
			// 	port = fmt.Sprintf("%d", layer.DstPort)
			// 	protocol = "UDP"
			default:
				continue
			}

			req := &snapwall.ServiceRequest{
				Time:        timestamppb.Now(),
				Type:        direction,
				Source:      srcIP,
				Destination: dstIP,
				Port:        port,
				Protocol:    protocol,
			}

			// isAllowed := true
			// for _, pol := range policies {
			// 	if !pol.Type {
			// 		isAllowed = matchPolicies(pol, req)
			// 	} else {
			// 		isAllowed = true
			// 	}
			// }

			// if !isAllowed {
			// 	log.Printf("SKIPPING PORT: %v, Source: %v, Destination: %v, Direction: %v", port, srcIP, dstIP, direction)
			// 	continue
			// }

			go func(req *snapwall.ServiceRequest) {
				stream, err := client.Send(context.Background())
				if err != nil {
					log.Printf("Failed to send packet data: %v", err)
					return
				}

				if err := stream.Send(req); err != nil {
					log.Printf("Failed to send request: %v", err)
					return
				}

				resp, err := stream.Recv()
				if err != nil {
					log.Printf("Failed to receive response: %v", err)
					return
				}

				fmt.Printf("Response from server: Time: %s | Source: %s | Destination: %s | Type: %s | Port: %s | Protocol: %s | Severity: %s\n",
					resp.Time, resp.Source, resp.Destination, resp.Type, resp.Port, resp.Protocol, resp.Severity)
			}(req)
		}
	}
}

func getLocalIPs() ([]string, error) {
	var localIPs []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && ip.To4() != nil {
				localIPs = append(localIPs, ip.String())
			}
		}
	}
	return localIPs, nil
}

func contains(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}
