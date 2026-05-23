package models

import "time"

type TargetType string

const (
	TypeNetworkTarget TargetType = "NETWORK_INFRASTRUCTURE"
	TypePhoneTarget   TargetType = "TELEPHONY_VECTOR"
	TypeEmailTarget   TargetType = "EMAIL_STEALTH_VECTOR"
	TypeGeoTarget     TargetType = "PRECISION_GEO_PING"
)

type IntelPayload struct {
	Target       string             `json:"target"`
	Type         TargetType         `json:"target_type"`
	ScanTime     time.Time          `json:"scan_time"`
	ASN          string             `json:"asn,omitempty"`
	ISP          string             `json:"isp,omitempty"`
	Geo          GeoData            `json:"geo,omitempty"`
	Clusters     []NamespaceCluster `json:"authoritative_clusters,omitempty"`
	Phone        PhoneData          `json:"phone_intel,omitempty"`
	Email        EmailData          `json:"email_intel,omitempty"`
	Verbose      bool               `json:"-"`
	OutputFormat string             `json:"-"`
}

type GeoData struct {
	Country      string `json:"country,omitempty"`
	Region       string `json:"region,omitempty"`
	RegionID     string `json:"region_id,omitempty"`
	City         string `json:"city,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	MapReference string `json:"map_reference,omitempty"`
}

type PhoneData struct {
	Valid       string `json:"valid"`
	LocalFormat string `json:"local_format"`
	CountryCode string `json:"country_code"`
	Location    string `json:"location"`
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
}

type EmailData struct {
	Deliverable string   `json:"deliverable"`
	Username    string   `json:"username"`
	Domain      string   `json:"domain"`
	MXRecords   []string `json:"mx_records,omitempty"`
	Disposable  string   `json:"disposable"`
	ProfileLink string   `json:"profile_link,omitempty"`
}

type NamespaceCluster struct {
	NameServer string   `json:"nameserver"`
	IPs        []string `json:"ips,omitempty"`
}
