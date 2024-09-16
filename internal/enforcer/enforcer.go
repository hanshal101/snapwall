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

				if policy.Type == "enforcer" {
					err = enforceRule(ipt, ip.Address, port.Number)
				} else if policy.Type == "deforcer" {
					err = deforceRule(ipt, ip.Address, port.Number)
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

func enforceRule(ipt *iptables.IPTables, ip, port string) error {
	log.Printf("Executing: iptables -A INPUT -s %s -p tcp --dport %s -j DROP\n", ip, port)

	err := ipt.AppendUnique("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to enforce rule for IP: %s, Port: %s, error: %v", ip, port, err)
	}

	log.Printf("Enforced rule for IP: %s, Port: %s\n", ip, port)
	return nil
}

func deforceRule(ipt *iptables.IPTables, ip, port string) error {
	log.Printf("Executing: iptables -D INPUT -s %s -p tcp --dport %s -j DROP\n", ip, port)

	err := ipt.Delete("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to deforce rule for IP: %s, Port: %s, error: %v", ip, port, err)
	}
	log.Printf("Deforced rule for IP: %s, Port: %s\n", ip, port)
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

				err = deforceRule(ipt, ip.Address, port.Number)
				if err != nil {
					log.Printf("Error processing policy %s: %v\n", policy.Name, err)
				} else {
					log.Printf("Processed policy %s (Type: %s) for IP: %s, Port: %s\n", policy.Name, policy.Type, ip.Address, port.Number)
				}

			}(policy, ip, port)
		}
	}

	wg.Wait()

	log.Println("Escaping Deletion !!!")
	return nil
}
