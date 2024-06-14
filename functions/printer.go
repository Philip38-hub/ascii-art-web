package asciiart

import (
	"strings"
)

func Printer(s string, b [][]string) string {
	var art strings.Builder
	// map each character of a string to its ascii character in our set
	for i := 0; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			art.WriteString(b[toPrint][i])
		}
		if i != 8 {
			art.WriteString("\n")
		}
	}
	return art.String()
}
