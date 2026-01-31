case "ping":
    // Attempting a connection-based ping (Standard for non-root)
    conn, err := net.DialTimeout("ip4:icmp", target, 2*time.Second)
    if err == nil {
        conn.Close()
        return true, time.Since(start)
    }
    // Fallback: Many environments block ICMP, so we check if the host is at least reachable
    _, err = net.LookupHost(target)
    return err == nil, time.Since(start)
