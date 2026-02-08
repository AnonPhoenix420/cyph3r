package output

import (
	"fmt"
	"time"
)

func PulseNode(target string) {
	fmt.Printf("%s[*] ESTABLISHING PULSE WITH: %s%s%s\n", White, NeonPink, target, Reset)
	time.Sleep(500 * time.Millisecond)
}
