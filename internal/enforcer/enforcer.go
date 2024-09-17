package enforcer

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/coreos/go-iptables/iptables"
	"github.com/hanshal101/snapwall/models"
)

func ReconcileEnforcer(
	ctx context.Context,
	policy models.Policy,
	ips []models.IP,
	ports []models.Port,
	application models.Application,
) error {

	log.Println("Reconciler Enforcer Started !!!")

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating iptables instance: %v\n", err)
		return err
	}

	var wg sync.WaitGroup

	log.Println("in loop !!!")
	for _, ip := range ips {
		for _, port := range ports {
			wg.Add(1)

			go func(policy models.Policy, ip models.IP, port models.Port) {
				defer wg.Done()

				log.Printf("Working for %v of %v from %v \n", policy.Type, ip.Address, port.Number)

				if policy.Type == "INGRESS" {
					err = INGRESS_RULE(ipt, ip.Address, application.Port, port.Number)
				} else if policy.Type == "EGRESS" {
					err = EGRESS_RULE(ipt, ip.Address, application.Port, port.Number)
				}

				if err != nil {
					log.Printf("Error processing policy %s: %v\n", policy.Name, err)
				} else {
					log.Printf("Processed policy %s (Type: %s) for IP: %s, Port: %s\n", policy.Name, policy.Type, ip.Address, port.Number)
				}
			}(policy, ip, port)
		}
	}

	wg.Wait()
	log.Println("Reconciler Enforcer Stopped !!!")
	return nil
}

func INGRESS_RULE(ipt *iptables.IPTables, srcIP, srcPort, dstPort string) error {
	log.Printf("Executing: iptables -A INPUT -s %s -p tcp --sport %s --dport %s -j DROP\n", srcIP, srcPort, dstPort)

	err := ipt.AppendUnique("filter", "INPUT", "-s", srcIP, "-p", "tcp", "--sport", srcPort, "--dport", dstPort, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to enforce incoming block from IP: %s Port: %s to Port: %s, error: %v", srcIP, srcPort, dstPort, err)
	}

	log.Printf("Enforced rule to block incoming requests from IP: %s Port: %s to Port: %s\n", srcIP, srcPort, dstPort)
	return nil
}

func EGRESS_RULE(ipt *iptables.IPTables, ip, srcPort, dstPort string) error {
	log.Printf("Executing: iptables -A OUTPUT -p tcp --sport %s -d %s --dport %s -j DROP\n", srcPort, ip, dstPort)

	err := ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "--sport", srcPort, "-d", ip, "--dport", dstPort, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to enforce outgoing block from Port: %s to IP: %s on Port: %s, error: %v", srcPort, ip, dstPort, err)
	}

	log.Printf("Enforced rule to block outgoing requests from Port: %s to IP: %s on Port: %s\n", srcPort, ip, dstPort)
	return nil
}

func RemoveINGRESSRule(ipt *iptables.IPTables, ip, port string) error {
	log.Printf("Executing: iptables -D INPUT -s %s -p tcp --dport %s -j DROP\n", ip, port)

	err := ipt.Delete("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to remove enforced rule for IP: %s, Port: %s, error: %v", ip, port, err)
	}
	log.Printf("Removed enforced rule for IP: %s, Port: %s\n", ip, port)
	return nil
}

func RemoveEGRESSRule(ipt *iptables.IPTables, ip, port string) error {
	log.Printf("Executing: iptables -D OUTPUT -d %s -p tcp --sport %s -j DROP\n", ip, port)

	err := ipt.Delete("filter", "OUTPUT", "-d", ip, "-p", "tcp", "--sport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to remove outgoing block for IP: %s, Port: %s, error: %v", ip, port, err)
	}

	log.Printf("Removed outgoing block for IP: %s, Port: %s\n", ip, port)
	return nil
}

func DeleteRule(
	ctx context.Context,
	policy models.Policy,
	ips []models.IP,
	ports []models.Port,
) error {
	log.Println("Delettion request made !!!")

	fmt.Println("del req made for:", policy)

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating iptables instance: %v\n", err)
		return err
	}

	var wg sync.WaitGroup

	log.Println("in loop !!!")
	for _, ip := range ips {
		for _, port := range ports {
			wg.Add(1)

			go func(policy models.Policy, ip models.IP, port models.Port) {
				defer wg.Done()
				if policy.Type == "INGRESS" {
					err = RemoveINGRESSRule(ipt, ip.Address, port.Number)
					if err != nil {
						log.Printf("Error processing policy %s: %v\n", policy.Name, err)
					}
				} else if policy.Type == "EGRESS" {
					err = RemoveEGRESSRule(ipt, ip.Address, port.Number)
					if err != nil {
						log.Printf("Error processing policy %s: %v\n", policy.Name, err)
					}
				} else {
					log.Fatalf("Error in removing Rules : %v\n", err)
					return
				}

				log.Printf("Processed policy %s (Type: %s) for IP: %s, Port: %s\n", policy.Name, policy.Type, ip.Address, port.Number)

			}(policy, ip, port)
		}
	}

	wg.Wait()

	log.Println("Escaping Deletion !!!")
	return nil
}
