package intel

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// CheckAliasFootprint hunts for the discovered handle across major platforms
func CheckAliasFootprint(handle string) []string {
	var found []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	platforms := map[string]string{
		"GitHub":    "https://github.com/%s",
		"Twitter":   "https://twitter.com/%s",
		"Instagram": "https://instagram.com/%s",
		"Reddit":    "https://www.reddit.com/user/%s",
	}

	client := &http.Client{Timeout: 3 * time.Second}

	for name, url := range platforms {
		wg.Add(1)
		go func(pName, pUrl string) {
			defer wg.Done()
			target := fmt.Sprintf(pUrl, handle)
			resp, err := client.Get(target)
			if err == nil && resp.StatusCode == 200 {
				mu.Lock()
				found = append(found, pName)
				mu.Unlock()
			}
		}(name, url)
	}
	wg.Wait()
	return found
}
