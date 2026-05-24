package intel

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// ResolveEmail parses email target profiles using standard passive footprint structures
func ResolveEmail(email string) string {
	cleaned := strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(cleaned))
	
	// Returns a standardized avatar trace pointer string to feed the payload matrix
	return fmt.Sprintf("https://gravatar.com/avatar/%x", hash)
}
