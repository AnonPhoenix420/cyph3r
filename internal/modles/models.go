package models

// IntelData is the central data structure for the Cyph3r system.
// It bridges the gap between intelligence gathering and the Neon HUD.
type IntelData struct {
	// --- Network Identity (Neon Blue) ---
	IPs         []string // List of resolved A records
	Nameservers []string // List of resolved NS records

	// --- Registry Intelligence (Neon Yellow) ---
	WhoisRaw    string   // The full, unparsed WHOIS response
	Registrar   string   // The domain registrar (e.g., GoDaddy, Namecheap)
	CreatedDate string   // Domain registration date
	ExpiryDate  string   // Domain expiration date

	// --- Geographic Intelligence (Neon Yellow) ---
	Country     string   // Country name
	CountryCode string   // ISO Country Code
	Region      string   // State/Province
	City        string   // City name
	Zip         string   // Postal/Zip code
	Timezone    string   // Target's local timezone
	Lat         float64  // Latitude for map link
	Lon         float64  // Longitude for map link

	// --- infrastructure Data (Neon Yellow) ---
	ISP         string   // Internet Service Provider
	Org         string   // Organization name
	ASN         string   // Autonomous System Number

	// --- Connectivity Status (Neon Green) ---
	IsAlive     bool     // Basic ping/handshake success
}

// ProbeResult tracks the outcome of individual port scans
type ProbeResult struct {
	Port     int      // The target port
	State    string   // ALIVE, CLOSED, or FILTERED
	Service  string   // Identified service (e.g., SSH, HTTP)
	Protocol string   // TCP/UDP
}
