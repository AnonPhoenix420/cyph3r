package output


import "fmt"


const (
Red = "\033[31m"
Blue = "\033[34m"
Reset = "\033[0m"
)


func Banner() {
fmt.Println("==============================")
fmt.Println(" cyph3r — Network Utility")
fmt.Println(" Educational use only ⚠️")
fmt.Println("==============================")
}


func Up(msg string) {
fmt.Println(Blue + "[UP] " + msg + Reset)
}


func Down(msg string) {
fmt.Println(Red + "[DOWN] " + msg + Reset)
}
