package asciiart

import (
	"errors"
	"os"
	"strings"
)

// Ascii generates ASCII art from input text using a specified format.
func Ascii(txt, format string) (string, error) {
	str := ""
	// Normalize line endings to '\n' if '\r\n' is found
	txt = strings.ReplaceAll(txt, "\r\n", "\n")
	textSlice := strings.Split(txt, "\n")

	// Validate characters in input text
	if !charValidation(txt) {
		return "", errors.New("error: invalid char")
	}

	// Read ASCII art template file
	file, err := os.ReadFile("ascii-art/banner/" + format + ".txt")
	if err != nil {
		return "", errors.New("error: reading file")
	}
	slice := strings.Split(string(file), "\n")

	// Generate ASCII art based on input text
	for j, txt := range textSlice {
		if txt != "" {
			// Generate 8 lines for each line of text
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					// Calculate index to fetch ASCII art line for each character
					firstLine := int(v-32)*9 + 1 + i
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n" // Add empty line between non-empty lines of text
		}
	}
	return str, nil
}

// charValidation checks if input string contains valid ASCII characters.
func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		// Valid characters are within ASCII range 32-126 or newline '\n'
		if (char < 32 || char > 126) && char != '\n' {
			return false
		}
	}
	return true
}
