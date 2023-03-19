package prompt

import (
	"fmt"

	"github.com/gookit/color"
)

func Warn(output string) {
	color.Warn.Println("⚠️ ", output)
}

func Success(output string) {
	color.Successln("✅ ", output)
}

func ClearScreen() {
	// Clears the terminal with an ANSI escape code.
	// Works in UNIX and newer Windows terminals.
	fmt.Println("\x1Bc")
}
