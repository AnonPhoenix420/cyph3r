package output

import "fmt"

func Info(msg string) {
	fmt.Println(BlueText("[INFO] " + msg))
}

func Success(msg string) {
	fmt.Println(GreenText("[OK] " + msg))
}

func Warn(msg string) {
	fmt.Println(YellowText("[WARN] " + msg))
}

func Down(msg string) {
	fmt.Println(RedText("[DOWN] " + msg))
}
