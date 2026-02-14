func DisplayHUD(data models.IntelData) {
	fmt.Printf("\n%s--- [ NODE_INTELLIGENCE_HUD ] ---%s\n", NeonPink, Reset)
	fmt.Printf("%s[*] Registrar: %s%s\n", White, NeonGreen, data.Registrar)
	fmt.Printf("%s[*] Location:  %s%s, %s\n", White, NeonGreen, data.City, data.Country)
	
	// Create a clickable Google Maps link for the terminal
	mapURL := fmt.Sprintf("https://www.google.com/maps?q=%f,%f", data.Lat, data.Lon)
	fmt.Printf("%s[*] Map Link:  %s%s%s\n", White, NeonBlue, mapURL, Reset)
}
