func PrintWhoisHUD(data models.IntelData) {
	fmt.Printf("\n%s──[ REGISTRY INTELLIGENCE ]──%s\n", White, Reset)
	if data.WhoisRaw != "" {
		// We print just the first few lines to keep the HUD clean
		fmt.Printf("%s WHOIS_DATA_SNIPPET:\n%s", White, NeonYellow)
		// Logic to truncate or print specific lines like "Registrar"
		fmt.Println(" [RECORDS_RETRIVED_SUCCESSFULLY]")
	} else {
		fmt.Printf("%s [!] WHOIS_QUERY_FAILED%s\n", NeonPink, Reset)
	}
}
