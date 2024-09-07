package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/coreos/go-iptables/iptables"
	"github.com/hanshal101/snapwall/database/psql"
	"github.com/hanshal101/snapwall/models"
	"github.com/joho/godotenv"
)

var (
	ctx       = context.Background()
	timeLimit = 1 * time.Second
)

func Reconciler(ctx context.Context, tmd time.Duration) {
	tmt := time.NewTicker(tmd)
	defer tmt.Stop()

	for range tmt.C {
		currentTag := fmt.Sprintf("reconcile-%d", time.Now().Unix())

		log.Println("Reconciler Started Successfully !!!")

		var policies []models.Policy
		if err := psql.DB.Preload("IPs").Preload("Ports").Find(&policies).Error; err != nil {
			log.Printf("Error in fetching policies: %v", err)
			return
		}

		ipt, err := iptables.New()
		if err != nil {
			log.Fatalf("Error creating iptables instance: %v\n", err)
		}

		var wg sync.WaitGroup
		for _, policy := range policies {
			for _, ip := range policy.IPs {
				for _, port := range policy.Ports {
					wg.Add(1)

					go func(policy models.Policy, ip models.IP, port models.Port) {
						defer wg.Done()

						log.Printf("Working for %v of %v from %v \n", policy.Type, ip.Address, port.Number)

						if policy.Type == "enforcer" {
							err = enforceRule(ipt, ip.Address, port.Number, currentTag)
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
		}

		wg.Wait()

		err = clearOldRules(ipt, currentTag)
		if err != nil {
			log.Printf("Error clearing old rules: %v", err)
		}

		log.Println("Reconciler Stopped !!!")
	}
}

func clearInputChainRules() {
	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating iptables instance: %v\n", err)
	}

	rules, err := ipt.List("filter", "INPUT")
	if err != nil {
		log.Fatalf("Error listing iptables rules: %v\n", err)
	}

	for _, rule := range rules {
		if strings.Contains(rule, "-j DROP") {
			parts := strings.Split(rule, " ")

			var srcIP, dport string
			for i, part := range parts {
				if part == "-s" && i+1 < len(parts) {
					srcIP = parts[i+1]
				} else if part == "--dport" && i+1 < len(parts) {
					dport = parts[i+1]
				}
			}

			if srcIP != "" && dport != "" {
				err = ipt.Delete("filter", "INPUT", "-s", srcIP, "-p", "tcp", "--dport", dport, "-j", "DROP")
				if err != nil {
					log.Printf("Error deleting rule with srcIP %s and port %s: %v\n", srcIP, dport, err)
				} else {
					log.Printf("Successfully deleted rule with srcIP %s and port %s\n", srcIP, dport)
				}
			}
		}
	}
}

func enforceRule(ipt *iptables.IPTables, ip, port, tag string) error {
	log.Printf("Adding rule: iptables -A INPUT -s %s -p tcp --dport %s -j DROP --comment %s\n", ip, port, tag)

	err := ipt.AppendUnique("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP", "-m", "comment", "--comment", tag)
	if err != nil {
		return fmt.Errorf("failed to enforce rule for IP: %s, Port: %s, error: %v", ip, port, err)
	}

	log.Printf("Enforced rule for IP: %s, Port: %s with tag %s\n", ip, port, tag)
	return nil
}

func deforceRule(ipt *iptables.IPTables, ip, port string) error {
	log.Printf("Attempting to delete rule: iptables -D INPUT -s %s -p tcp --dport %s -j DROP\n", ip, port)

	exists, err := ipt.Exists("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to check if rule exists for IP: %s, Port: %s, error: %v", ip, port, err)
	}

	if !exists {
		log.Printf("Rule does not exist for IP: %s, Port: %s, nothing to delete", ip, port)
		return nil
	}

	err = ipt.Delete("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
	if err != nil {
		return fmt.Errorf("failed to delete rule for IP: %s, Port: %s, error: %v", ip, port, err)
	}

	log.Printf("Successfully deleted rule for IP: %s, Port: %s", ip, port)
	return nil
}

func normalizeIP(ip string) string {
	if strings.Contains(ip, "/32") {
		return strings.Replace(ip, "/32", "", -1)
	}
	return ip
}

func clearOldRules(ipt *iptables.IPTables, currentTag string) error {
	rules, err := ipt.List("filter", "INPUT")
	if err != nil {
		return fmt.Errorf("failed to list iptables rules: %v", err)
	}

	for _, rule := range rules {
		if !strings.Contains(rule, currentTag) && strings.Contains(rule, "reconcile-") {
			parts := strings.Split(rule, " ")
			var srcIP, dport string
			for i, part := range parts {
				if part == "-s" && i+1 < len(parts) {
					srcIP = normalizeIP(parts[i+1])
				} else if part == "--dport" && i+1 < len(parts) {
					dport = parts[i+1]
				}
			}

			if srcIP != "" && dport != "" {
				err = ipt.Delete("filter", "INPUT", "-s", srcIP, "-p", "tcp", "--dport", dport, "-j", "DROP")
				if err != nil {
					if strings.Contains(err.Error(), "Bad rule") {
						log.Printf("No matching rule found for deletion: srcIP %s, port %s\n", srcIP, dport)
					} else {
						log.Printf("Error deleting rule with srcIP %s and port %s: %v\n", srcIP, dport, err)
					}
				} else {
					log.Printf("Successfully deleted old rule with srcIP %s and port %s\n", srcIP, dport)
				}
			}
		}
	}

	return nil
}

func ruleExists(ipt *iptables.IPTables, ip, port string) (bool, error) {
	rules, err := ipt.List("filter", "INPUT")
	if err != nil {
		return false, fmt.Errorf("failed to list iptables rules: %v", err)
	}

	ruleToFind := fmt.Sprintf("-s %s -p tcp --dport %s -j DROP", ip, port)
	for _, rule := range rules {
		if strings.Contains(rule, ruleToFind) {
			return true, nil
		}
	}
	return false, nil
}

func addRuleIfNotExists(ipt *iptables.IPTables, ip, port string) error {
	exists, err := ruleExists(ipt, ip, port)
	if err != nil {
		return fmt.Errorf("error checking rule existence: %v", err)
	}
	if !exists {
		err := ipt.Append("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
		if err != nil {
			return fmt.Errorf("failed to add rule: %v", err)
		}
	}
	return nil
}

func deleteAllMatchingRules(ipt *iptables.IPTables, ip, port string) error {
	rules, err := ipt.List("filter", "INPUT")
	if err != nil {
		return fmt.Errorf("failed to list iptables rules: %v", err)
	}

	ruleToDelete := fmt.Sprintf("-s %s -p tcp --dport %s -j DROP", ip, port)
	for _, rule := range rules {
		if strings.Contains(rule, ruleToDelete) {
			err = ipt.Delete("filter", "INPUT", "-s", ip, "-p", "tcp", "--dport", port, "-j", "DROP")
			if err != nil {
				return fmt.Errorf("failed to delete rule: %v", err)
			}
		}
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	psql.InitDB()
	tickerDuration := 5 * time.Second
	go Reconciler(ctx, tickerDuration)

	select {}
}
