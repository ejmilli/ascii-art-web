package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GenerateASCIIArt generates ASCII art for the given text and template.
func GenerateASCIIArt(text, template string) (string, error) {
	templates := map[string]string{
		"standard":   "./ascii/txt/standard.txt",
		"shadow":     "./ascii/txt/shadow.txt",
		"thinkertoy": "./ascii/txt/thinkertoy.txt",
	}

	templatePath, exists := templates[template]
	if !exists {
		return "", fmt.Errorf("template not found: %s", template)
	}

	asciiMap := LoadTemplate(templatePath)
	if asciiMap == nil {
		return "", fmt.Errorf("failed to load template: %s", templatePath)
	}

	return RenderASCII(asciiMap, text), nil
}

// RenderASCII generates the ASCII art string.
func RenderASCII(asciiMap map[rune][]string, text string) string {
	var result strings.Builder
	lines := SplitTextByLines(ParseEscapeSequences(text))

	for _, line := range lines {
		for i := 0; i < 8; i++ {
			for _, char := range line {
				asciiLines, exists := asciiMap[char]
				if !exists {
					result.WriteString(fmt.Sprintf("Error: '%c' not found\n", char))
					return result.String()
				}
				result.WriteString(asciiLines[i])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}

// LoadTemplate and other helper functions remain the same
func LoadTemplate(filePath string) map[rune][]string {
	// Adjust the path to load templates from the 'ascii/templates' folder
	file, err := os.Open("ascii/templates/" + filePath) // Path adjusted
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
