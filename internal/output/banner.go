package output

import "fmt"

func Banner() {
	// I'm using your Neon Blue here. 
	// PASTE YOUR ORIGINAL ASCII ART BETWEEN THE BACKTICKS BELOW
	fmt.Printf("%s", NeonBlue)
	fmt.Println(`
   ______      ____  __  __ _____ ____  
  / ____/_  __/ __ \/ / / /|__  // __ \ 
 / /   / / / / /_/ / /_/ /  /_ </ /_/ / 
/ /___/ /_/ / ____/ __  / ___/ / _, _/  
\____/\__, /_/   /_/ /_/ /____/_/ |_|   
     /____/         NETWORK_INTEL_SYSTEM`)
	fmt.Printf("%s", Reset)
}
