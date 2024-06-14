package asciiart

import (
	"bufio"
	"os"
	"strings"
)

func CreateCharacters(bannername string, input string) string {
	file, _ := os.Open(bannername + ".txt")
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// group lines read into 9 to display each character individually
	character := make([]string, 9)
	var result [][]string
	for i := 1; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		character = lines[i:end]
		result = append(result, character)
	}

	// r := os.Args[1]
	// Format ("\r\n") in input string from the browser
	s := strings.Replace(input, "\r\n", "\n", -1) 
	s = strings.Replace(s, "\t", "    ", -1)

	return HandleLn(s, result)
}
