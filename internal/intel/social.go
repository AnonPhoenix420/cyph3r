package intel

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/AnonPhoenix420/cyph3r/internal/models"
)

// ResolveEmail parses email target profiles using standard passive footprint structures
func ResolveEmail(email string) string {
	cleaned := strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(cleaned))
	
	// Returns a standardized avatar trace pointer string to feed the payload matrix
	return fmt.Sprintf("https://gravatar.com/avatar/%x", hash)
}

// GetSocialProfiles returns elite-level social associations across multiple platforms
func GetSocialProfiles(target string, tType models.TargetType) []models.SocialProfile {
	if tType != models.TargetEmail && tType != models.TypeEmailTarget {
		return []models.SocialProfile{}
	}

	username := strings.Split(target, "@")[0]
	domain := strings.Split(target, "@")[1]

	profiles := []models.SocialProfile{
		// Gravatar (High confidence)
		{
			Platform:    "Gravatar",
			Username:    username,
			ProfileURL:  ResolveEmail(target),
			DisplayName: username,
			Confidence:  95,
		},
		// Google / Gmail
		{
			Platform:    "Google / Gmail",
			Username:    username,
			ProfileURL:  "https://myaccount.google.com",
			DisplayName: username,
			Confidence:  85,
		},
		// X (Twitter)
		{
			Platform:    "X (Twitter)",
			Username:    username,
			ProfileURL:  "https://x.com/" + username,
			Confidence:  50,
		},
		// LinkedIn
		{
			Platform:    "LinkedIn",
			Username:    username,
			ProfileURL:  "https://linkedin.com/in/" + username,
			Confidence:  45,
		},
		// Facebook
		{
			Platform:    "Facebook",
			Username:    username,
			ProfileURL:  "https://facebook.com/" + username,
			Confidence:  40,
		},
		// Reddit
		{
			Platform:    "Reddit",
			Username:    username,
			ProfileURL:  "https://reddit.com/u/" + username,
			Confidence:  40,
		},
		// WhatsApp (indirect via number or web)
		{
			Platform:    "WhatsApp",
			Username:    username,
			ProfileURL:  "https://wa.me/" + username, // Usually phone-based
			Confidence:  25,
		},
		// MySpace (legacy)
		{
			Platform:    "MySpace",
			Username:    username,
			ProfileURL:  "https://myspace.com/" + username,
			Confidence:  20,
		},
		// Additional common platforms
		{
			Platform:    "GitHub",
			Username:    username,
			ProfileURL:  "https://github.com/" + username,
			Confidence:  55,
		},
		{
			Platform:    "Instagram",
			Username:    username,
			ProfileURL:  "https://instagram.com/" + username,
			Confidence:  45,
		},
		{
			Platform:    "TikTok",
			Username:    username,
			ProfileURL:  "https://tiktok.com/@" + username,
			Confidence:  35,
		},
		{
			Platform:    "Pinterest",
			Username:    username,
			ProfileURL:  "https://pinterest.com/" + username,
			Confidence:  30,
		},
	}

	// Domain-specific hints
	if strings.Contains(domain, "gmail") || strings.Contains(domain, "google") {
		profiles = append(profiles, models.SocialProfile{
			Platform:    "Google Ecosystem",
			Username:    username,
			ProfileURL:  "https://plus.google.com",
			Confidence:  70,
		})
	}

	return profiles
}
