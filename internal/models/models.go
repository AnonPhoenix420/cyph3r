package models

import "time"

type IntelPayload struct {
	Target       string             `json:"target"`
	ScanTime     time.Time          `json:"scan_time"`
	ASN          string             `json:"asn,omitempty"`
	ISP          string             `json:"isp,omitempty"`
	Geo          GeoData            `json:"geo,omitempty"`
	Clusters     []NamespaceCluster `json:"authoritative_clusters,omitempty"`
	Verbose      bool               `json:"-"`
	OutputFormat string             `json:"-"`
}

type GeoData struct {
	Country  string `json:"country,omitempty"`
	Region   string `json:"region,omitempty"`
	RegionID string `json:"region_id,omitempty"`
	City     string `json:"city,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type NamespaceCluster struct {
	NameServer string   `json:"nameserver"`
	IPs        []string `json:"ips,omitempty"`
}
