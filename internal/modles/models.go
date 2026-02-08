package models

// IntelData is the master structure that holds all gathered intelligence.
// This is used by the HUD (render.go) to display your neon-colored results.
type IntelData struct {
	// Identity Data (Neon Blue & Neon Yellow)
	IPs         []string // Resolved IP addresses
	Nameservers []string // NS Node records

	// Geographic Data (Neon Yellow)
	Country     string  // Full country name
	CountryCode string  // ISO Code (e.g., US, UK)
	Region      string  // State or Province
	City        string  // City name
	Zip         string  // Postal code
	Timezone    string  // Local timezone of the target

	// Technical Metadata (Neon Green)
	ISP string // Internet Service Provider
	Org string // Organization name
	ASN string // Autonomous System Number

	// Coordinate Data (For the Neon Pink Map Link)
	Lat float64 // Latitude
	Lon float64 // Longitude
}

// ProbeResult holds the status of an individual port check.
type ProbeResult struct {
	Port     int    // Port number (e.g., 80)
	State    string // ALIVE, CLOSED, or FILTERED
	Service  string // Detected service (e.g., HTTP)
	Protocol string // TCP or UDP
}
