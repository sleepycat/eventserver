package scan

import (
	"context"
	"github.com/Ullaakut/nmap/v2"
	"log"
	"time"
)

// Port : a structure describing a port
type Port struct {
	Status      string `json:"status"`
	Number      int    `json:"number"`
	Description string `json:"description"`
}

// Scan : Given a domain/IP scans returns a list of open ports
func Scan(url string) []Port {
	results := []Port{}
	// scan localhost with a 2 second timeout per port in 5 concurrent threads
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(url),
		nmap.WithPorts("1-65535"),
		nmap.WithContext(ctx),
	)

	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			p := Port{Number: int(port.ID)}
			// p.Status = string(port.State)
			p.Description = port.Service.Name
			results = append(results, p)
		}
	}

	return results
}
