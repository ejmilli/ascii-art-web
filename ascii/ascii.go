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

	asciiMap, err := LoadTemplate(templatePath)

	if err != nil {
		// Return 500 status code for internal server error
		return "", http.StatusInternalServerError
	}

	asciiArt := RenderASCII(asciiMap, text)
	return asciiArt, 200
}

// RenderASCII generates the ASCII art string.
func RenderASCII(asciiMap map[rune][]string, text string) string {
	var result strings.Builder
	asciiArtLines := make([]string, 8) // Prepare a common array for the entire text's ASCII art

	for _, char := range text {
			asciiLines := asciiMap[char]
			// Append ASCII lines for this character to the corresponding line in the result
			for i := 0; i < 8; i++ {
					asciiArtLines[i] += asciiLines[i]
			}
	}

	// Add the complete ASCII lines for the entire text
	for _, asciiLine := range asciiArtLines {
			result.WriteString(asciiLine + "\n")
	}
	return result.String()
}

// LoadTemplate and other helper functions remain the same
func LoadTemplate(filePath string) (map[rune][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
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

	if len(lines) > 0 {
		asciiMap[character] = lines
	}
	return asciiMap, nil

}
