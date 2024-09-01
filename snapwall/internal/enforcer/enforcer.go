package enforcer

import (
	"context"
	"log"

	"github.com/hanshal101/snapwall/database"
	snapwall "github.com/hanshal101/snapwall/proto"
)

type Policy struct {
	Name  string
	Type  string
	IPs   []IP
	Ports []Port
}

type IP struct {
	PolicyID uint
	Address  string
}

type Port struct {
	PolicyID uint
	Number   string
}

func CheckPacket(ctx context.Context, pkt *snapwall.ServiceRequest) bool {
	var policies []Policy
	if err := database.DB.Find(&policies).Error; err != nil {
		log.Fatalf("Error in finding policies: %v\n", err)
		return false
	}
	for _, policy := range policies {
		isMatched := matchpolicies(policy, pkt)
		if !isMatched {
			log.Fatalf("Intrusion Found by %v!!!\nPort: %v, Source: %v, Destination: %v, Protocol: %v\n",
				policy.Name, pkt.Port, pkt.Source, pkt.Destination, pkt.Protocol)
			return false
		}
	}

	return true
}

func matchpolicies(policy Policy, req *snapwall.ServiceRequest) bool {
	for _, pol := range policy.IPs {
		if req.Source == pol.Address {
			return false
		}
	}
	for _, pol := range policy.Ports {
		if req.Port == pol.Number {
			return false
		}
	}
	return true
}

func enforcePKT(req *snapwall.ServiceRequest) error {
	
	return nil
}
