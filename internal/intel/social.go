package intel

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func CheckAliasFootprint(handle string) []string {
	var found []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	platforms := map[string]string{
		"GitHub": "https://github.com/%s",
		"Reddit": "https://www.reddit.com/user/%s",
		"X":      "https://twitter.com/%s",
	}
	client := &http.Client{Timeout: 3 * time.Second}
	for name, url := range platforms {
		wg.Add(1)
		go func(n, u string) {
			defer wg.Done()
			resp, err := client.Head(fmt.Sprintf(u, handle))
			if err == nil && resp.StatusCode == 200 {
				mu.Lock(); found = append(found, n); mu.Unlock()
			}
		}(name, url)
	}
	wg.Wait()
	return found
}
