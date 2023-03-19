package prompt

import (
	"github.com/gookit/color"
)

func Warn(output string) {
	color.Warn.Println("⚠️ ", output)
}

func Success(output string) {
	color.Successln("✅ ", output)
}
