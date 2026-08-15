package stress

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ExecuteContinuousStress(targetAddr string, concurrency int, durationSec int) {
	fmt.Printf("[!] ENGAGING CONTINUOUS HULK ENGINE: %s (Concurrency: %d)\n", targetAddr, concurrency)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	taskStream := make(chan struct{}, concurrency)

	// Spawn persistent worker pool
	for i := 0; i < concurrency; i++ {
		go func() {
			dialer := &net.Dialer{Timeout: 2 * time.Second}
			for range taskStream {
				conn, err := dialer.Dial("tcp", targetAddr)
				if err == nil {
					_ = conn.SetDeadline(time.Now().Add(1 * time.Second))
					conn.Close()
				}
			}
		}()
	}

	// Task Feeder Routine
	go func() {
		for {
			taskStream <- struct{}{}
		}
	}()

	if durationSec > 0 {
		timer := time.NewTimer(time.Duration(durationSec) * time.Second)
		select {
		case <-timer.C:
			fmt.Println("\n[+] Resilience test duration completed.")
		case <-sigChan:
			fmt.Println("\n[+] Resilience test aborted by operator.")
		}
	} else {
		// Infinite duration mode (-d 0)
		<-sigChan
		fmt.Println("\n[+] Continuous stress loop terminated by operator.")
	}
}
