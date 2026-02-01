package output

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	CyanText   = color.New(color.FgCyan).SprintFunc()
	WhiteText  = color.New(color.FgWhite, color.Bold).SprintFunc()
	YellowText = color.New(color.FgYellow).SprintFunc()
	BlueText   = color.New(color.FgBlue).SprintFunc()
	RedText    = color.New(color.FgRed).SprintFunc()
	GreenText  = color.New(color.FgGreen).SprintFunc()
	MagText    = color.New(color.FgMagenta).SprintFunc()
)

func Info(m string) {
	fmt.Printf("%s %s\n", BlueText("[*]"), m)
}
