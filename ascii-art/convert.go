package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide a valid text argument. Usage: go run. <text> <template>")
		return
	}

	// Parse escape sequences (like "\n") in the provided text argument.
	// This function replaces any occurrences of `\n` with actual newline characters.
	text := os.Args[1]
	template := os.Args[2]

	templates := map[string]string{
		"standard":   "standard.txt",
		"shadow":     "shadow.txt",
		"thinkertoy": "thinkertoy.txt",
	}
	templatePath, exists := templates[template]
	if !exists {
		fmt.Println("Template not found.")
		return
	}
	// Load the ASCII art template from a file called "standard.txt".
	// This function reads the file, stores ASCII representations of characters in a map, and returns the map.
	asciiMap := LoadTemplate(templatePath)

	PrintASCII(asciiMap, text)
}

func PrintASCII(asciiMap map[rune][]string, text string) {

	lines := SplitTextByLines(ParseEscapeSequences(text))

	chars := 0
	for _, line := range lines {
		chars += len(line)
	}
	if chars == 0 {
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {

		for i := 0; i < 8; i++ {
			for _, char := range line {
				asciiLines, exists := asciiMap[char]
				if !exists {
					fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
					return
				}
				fmt.Print(asciiLines[i])
			}
			fmt.Println()
		}
	}
}

func LoadTemplate(filePath string) map[rune][]string {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	asciiMap := make(map[rune][]string)

	scanner := bufio.NewScanner(file)

	var character rune = ' '

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(lines) > 0 {
				asciiMap[character] = lines
				character++
				lines = []string{}
			}
		} else {

			lines = append(lines, line)
		}
	}

	return asciiMap
}

func ParseEscapeSequences(input string) string {
	return strings.ReplaceAll(input, `\n`, "\n")
}

func SplitTextByLines(text string) []string {
	return strings.Split(text, "\n")
}
