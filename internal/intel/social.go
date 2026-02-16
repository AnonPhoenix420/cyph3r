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
		"GitHub":    "https://github.com/%s",
		"Reddit":    "https://www.reddit.com/user/%s",
		"Twitter":   "https://twitter.com/%s",
		"Instagram": "https://instagram.com/%s",
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
		// Prevent following redirects to login pages which can give false positives
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for name, urlTemplate := range platforms {
		wg.Add(1)
		go func(pName, pUrl string) {
			defer wg.Done()
			target := fmt.Sprintf(pUrl, handle)
			resp, err := client.Head(target) 
			if err == nil && resp.StatusCode == 200 {
				mu.Lock()
				found = append(found, pName)
				mu.Unlock()
			}
		}(name, urlTemplate)
	}
	wg.Wait()
	return found
}
