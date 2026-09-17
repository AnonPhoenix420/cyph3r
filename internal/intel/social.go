package intel

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ResolveSocialFootprint sweeps globally recognized surface and deep-web systems for target identities
func ResolveSocialFootprint(username string) []models.SocialProfile {
	cleaned := strings.TrimSpace(username)
	var profiles []models.SocialProfile

	// Complete global and legacy footprint platform matrix configuration mapping
	platforms := []struct {
		Name string
		Base string
	}{
		{"X / Twitter", "https://x.com/"},
		{"Facebook", "https://www.facebook.com/"},
		{"LinkedIn", "https://www.linkedin.com/in/"},
		{"Myspace", "https://myspace.com/"},
		{"Pinterest", "https://www.pinterest.com/"},
		{"Reddit", "https://www.reddit.com/user/"},
		{"Instagram", "https://www.instagram.com/"},
		{"TikTok", "https://www.tiktok.com/@"},
		{"Tor2Web Onion Index", "https://ahmia.fi/search/?q="}, // Deep/Dark Web search routing path
	}

	client := &http.Client{
		Timeout: 4 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 3 {
				return fmt.Errorf("stopped after 3 redirects")
			}
			return nil
		},
	}

	for _, plat := range platforms {
		profileURL := fmt.Sprintf("%s%s", plat.Base, cleaned)
		confidenceScore := 75
		if len(cleaned) >= 6 {
			confidenceScore = 85
		}
		
		bio := fmt.Sprintf("Active fingerprint tracking verification pending deep socket parsing on %s.", plat.Name)
		statusTag := "UNVERIFIED"

		// Perform active HTTP probe for surface platforms (exclude search engines like Ahmia)
		if !strings.Contains(plat.Base, "ahmia.fi") {
			req, err := http.NewRequest("GET", profileURL, nil)
			if err == nil {
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Cyph3rOSINT/1.0")
				resp, err := client.Do(req)
				if err == nil {
					defer resp.Body.Close()
					if resp.StatusCode == http.StatusOK {
						confidenceScore = 95
						statusTag = "CONFIRMED_ACTIVE"
						bio = fmt.Sprintf("Target identity validated via HTTP 200 response on %s.", plat.Name)
					} else if resp.StatusCode == http.StatusNotFound {
						confidenceScore = 20
						statusTag = "NOT_FOUND"
						bio = fmt.Sprintf("Identity node returned 404 Not Found on %s.", plat.Name)
					} else {
						bio = fmt.Sprintf("Identity node returned HTTP status code %d on %s.", resp.StatusCode, plat.Name)
					}
				}
			}
		} else {
			statusTag = "INDEX_QUERY"
			bio = fmt.Sprintf("Deep web index query constructed for identity token: %s.", cleaned)
		}

		profiles = append(profiles, models.SocialProfile{
			Platform:    plat.Name,
			Username:    cleaned,
			ProfileURL:  profileURL,
			DisplayName: fmt.Sprintf("Identity Node ➔ %s [%s]", cleaned, statusTag),
			Bio:         bio,
			Confidence:  confidenceScore,
		})
	}

	return profiles
}
