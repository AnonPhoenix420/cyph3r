func DisplayPhoneHUD(p models.PhoneData) {
	fmt.Printf("\n%s--- [ 🛰️ GLOBAL_SATELLITE_HUD ] ---\n", NeonPink)
	fmt.Printf("%s[*] Target:     %s%s\n", White, NeonBlue, p.Number)
	fmt.Printf("%s[*] Status:     %s%t\n", White, NeonGreen, p.Valid)
	fmt.Printf("%s[*] Type:       %s%s\n", White, NeonYellow, p.Type)
	fmt.Printf("%s[*] Risk Level: %s%s\n", White, NeonPink, p.Risk)
	fmt.Printf("%s[*] Carrier:    %s%s\n", White, NeonYellow, p.Carrier)
	fmt.Printf("%s[*] Location:   %s%s, %s\n", White, NeonGreen, p.Location, p.Country)
	fmt.Printf("%s[*] Map Vector: %s%s\n", White, NeonBlue, p.MapLink)
	fmt.Printf("%s------------------------------------%s\n", NeonPink, Reset)
}
