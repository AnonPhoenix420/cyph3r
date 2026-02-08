package output

import "fmt"

// PrintNodeIntelligence formats the target and all its discovered addresses
func PrintNodeIntelligence(target string, ips []string, ns []string) {
    fmt.Printf("%s──[ NODE INTELLIGENCE ]──%s\n", White, Reset)
    fmt.Printf(" TARGET: %s\n", target)
    
    // Join IPs with a clean separator
    fmt.Print(" IP ADDRESSES: ")
    for i, ip := range ips {
        fmt.Print(CyanText(ip))
        if i < len(ips)-1 { fmt.Print(" | ") }
    }
    fmt.Println()

    // Render NS records in a clean list
    for _, node := range ns {
        fmt.Printf(" NS NODE:    %s\n", YellowText(node))
    }
}
