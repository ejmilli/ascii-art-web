package ascii

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func CleanInput(text string) string {
	return strings.ReplaceAll(text, "\r", "")
}

// GenerateASCIIArt generates ASCII art for the given text and template.
func GenerateASCIIArt(text, template string) (string, int) {

	text = CleanInput(text)
	// Validate input for unprintable ASCII characters
	for _, char := range text {
		if char < 32 || char > 126 {
			if char != '\n' && char != ' ' { // Allow spaces and newlines
				return "", http.StatusBadRequest
			}
		}
	}

	templates := map[string]string{
		"standard":   "./ascii/txt/standard.txt",
		"shadow":     "./ascii/txt/shadow.txt",
		"thinkertoy": "./ascii/txt/thinkertoy.txt",
	}

	templatePath := templates[template]

	asciiMap := LoadTemplate(templatePath)

	asciiArt := RenderASCII(asciiMap, text)
	return asciiArt, 200
}

// RenderASCII generates the ASCII art string.
func RenderASCII(asciiMap map[rune][]string, text string) string {
	var result strings.Builder
	lines := SplitTextByLines(ParseEscapeSequences(text))

	for _, line := range lines {
		// Prepare an array for each line of ASCII art (8 lines)
		asciiArtLines := make([]string, 8)

		for _, char := range line {
			asciiLines := asciiMap[char]
			// Append ASCII lines for this character to the corresponding line in the result
			for i := 0; i < 8; i++ {
				asciiArtLines[i] += asciiLines[i]
			}
		}

		// Add the complete ASCII lines for the current input line
		for _, asciiLine := range asciiArtLines {
			result.WriteString(asciiLine + "\n")
		}
	}
	return result.String()
}

// LoadTemplate and other helper functions remain the same
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
				lines = nil
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
