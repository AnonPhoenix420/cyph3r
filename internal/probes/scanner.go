package probes

import (
	"sync"
	"github.com/AnonPhoenix420/cyph3r/internal/output"
)

// RunFullScan orchestrates the multi-threaded probe sequence.
func RunFullScan(target string) {
	// Define our target tech stack ports
	ports := []int{
		21, 22, 23, 25, 53, 80, 110, 135, 139, 
		143, 443, 445, 993, 995, 1723, 3306, 
		3389, 5900, 8080, 8443,
	}

	output.PrintScanHeader()

	// WaitGroup to ensure we wait for all "Waves" to return
	var wg sync.WaitGroup
	
	// Semaphore to limit concurrency (don't overwhelm the local network)
	maxGuards := make(chan struct{}, 10) 

	for _, port := range ports {
		wg.Add(1)
		maxGuards <- struct{}{} // Occupy a slot

		go func(p int) {
			defer wg.Done()
			defer func() { <-maxGuards }() // Release slot

			// Call the logic from your probes.go
			_, status, _ := ConductWave(target, p)

			// Only report if the port is ALIVE to keep the HUD clean, 
			// or you can remove the 'if' to show everything.
			if status == "ALIVE" {
				output.PrintWaveStatus(p, status)
			}
		}(port)
	}

	wg.Wait()
}
