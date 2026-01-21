package output

import "fmt"

func Info(msg string) {
	fmt.Println(BlueText("[INFO] " + msg))
}

func Success(msg string) {
	fmt.Println(GreenText("[SUCCESS] " + msg))
}

func Down(msg string) {
	fmt.Println(RedText("[DOWN] " + msg))
}
