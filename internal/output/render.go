package output

import "fmt"

// PrintGeoHUD renders the geo-intelligence with a visual map link
func PrintGeoHUD(city, country, lat, lon string) {
	// Google Maps URL format for direct coordinate pin
	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%s,%s", lat, lon)

	fmt.Printf("\n%s[🛰️ DEEP NODE INTELLIGENCE]%s\n", Cyan, Reset)
	fmt.Printf("📍 Location: %s, %s\n", city, country)
	fmt.Printf("🗺️  Map Link: %s\n", mapURL)
	fmt.Println("------------------------------------------------")
}

// PrintWaveStatus displays the multi-probe results in a table format
func PrintWaveStatus(port int, method, status, conversation string) {
	fmt.Printf("%sPORT %-5d%s | %-8s | %-10s | %s\n", 
		Blue, port, Reset, method, status, conversation)
}
