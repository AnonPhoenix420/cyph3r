package models

// IntelData is the master structure for the Cyph3r intelligence HUD.
// It aggregates network, registry, and geographic metadata.
type IntelData struct {
	// --- Network Identity ---
	IPs         []string // Resolved IPv4/IPv6 addresses
	Nameservers []string // Target's authoritative nameservers

	// --- Registry Intelligence ---
	WhoisRaw    string   // Full raw WHOIS response for deep inspection
	Registrar   string   // Domain registrar name (e.g., Namecheap, Cloudflare)
	CreatedDate string   // ISO date of registration
	ExpiryDate  string   // ISO date of expiration

	// --- Geographic Intelligence ---
	Country     string   // Physical country of the target IP
	CountryCode string   // ISO-2 (e.g., US, UK, DE)
	Region      string   // State or Province
	City        string   // Closest metropolitan city
	Zip         string   // Postal code
	Timezone    string   // Local time zone of the target
	Lat         float64  // Latitude coordinate
	Lon         float64  // Longitude coordinate

	// --- Infrastructure Data ---
	ISP         string   // Internet Service Provider (e.g., AWS, DigitalOcean)
	Org         string   // Registered organization name
	ASN         string   // Autonomous System Number (e.g., AS15169)

	// --- Connectivity Status ---
	IsAlive     bool     // Boolean check for ping/handshake success
}

// ProbeResult tracks the status of specific port scans.
type ProbeResult struct {
	Port     int      // Port number (e.g., 80, 443, 22)
	State    string   // Status string: ALIVE, CLOSED, or FILTERED
	Service  string   // Guessed service (e.g., HTTP, SSH, MySQL)
	Protocol string   // Network protocol (TCP/UDP)
}
