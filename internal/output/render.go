// Inside DisplayHUD in render.go, update the port loop:

if ports := data.NameServers["PORTS"]; len(ports) > 0 {
    fmt.Printf("\n%s[*] INFO: Initializing Tactical Admin Scan: %s%s\n", White, NeonPink, data.TargetName)
    for _, p := range ports {
        if strings.Contains(p, "[!]") {
            // Highlight vulnerabilities in Red/Yellow
            fmt.Printf("%s[!] ALERT: PORT %s: %sCRITICAL_VERSION_DETECTED\n", NeonYellow, p, White)
        } else {
            fmt.Printf("%s[+] PORT %s: %sOPEN [ACK/SYN]\n", NeonGreen, p, White)
        }
    }
    fmt.Printf("%s[*] INFO: Admin/Web scan complete.\n", White)
}
