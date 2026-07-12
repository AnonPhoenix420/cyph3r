// Add this inside your main() function's execution logic:

// ... after flag.Parse() ...

// Resilience Test Routing
if *hulkFlag {
    targetURL := "http://" + cleanHost
    targetAddr := net.JoinHostPort(cleanHost, fmt.Sprintf("%d", *portFlag))

    switch strings.ToLower(*protoFlag) {
    case "udp":
        stress.ExecuteUDPFlood(targetAddr, *concurrencyFlag, *durationFlag)
    case "tcp":
        stress.ExecuteTCPFlood(targetAddr, *concurrencyFlag, *durationFlag)
    case "http":
        stress.ExecuteHTTPCapacityTest(targetURL, strings.ToUpper(*methodFlag), *concurrencyFlag, *durationFlag)
    default:
        fmt.Println("[-] Protocol not supported for stress testing. Use --proto tcp/udp/http")
    }
    return
}
