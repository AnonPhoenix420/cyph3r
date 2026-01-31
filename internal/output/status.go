package output

import "fmt"

func Info(msg string)    { fmt.Printf("%s[*] %s%s\n", Cyan, msg, Reset) }
func Success(msg string) { fmt.Printf("%s[+] %s%s\n", Green, msg, Reset) }
func Warn(msg string)    { fmt.Printf("%s[!] %s%s\n", Yellow, msg, Reset) }
func Down(msg string)    { fmt.Printf("%s[-] %s%s\n", Red, msg, Reset) }
