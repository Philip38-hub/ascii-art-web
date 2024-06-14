package asciiart

import (
	"os"
	"strings"
)

func HandleLn(s string, b [][]string) string{
	var printedLine strings.Builder
	if s == "" {
		printedLine.WriteString("\n")
		os.Exit(0)
	}
	//Check if all characters in s are exclusively "\n"
	isAllNewline := true
	for _, char := range s {
		if char != '\n' {
			isAllNewline = false
			break
		}
	}

	if isAllNewline {
		count := strings.Count(s, "\n")
		for i := 0; i < count; i++ {
			printedLine.WriteString("\n")
		}
		os.Exit(0)
	}
	//logic for processing lines
	str := strings.Split(s, "\n")
	for _, char := range str {
		if char == "" {
			printedLine.WriteString("\n")
			continue
		}
		printedLine.WriteString(Printer(char, b))
	}
	return printedLine.String()
}
