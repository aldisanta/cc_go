package kata

import (
	"bytes"
	"strings"
)

// DecodeMorse to decode morse code
func DecodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)
	var decoded bytes.Buffer
	i := 0
	for _, words := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		if i > 0 {
			decoded.WriteString(" ")
		}
		i++
		for _, letter := range strings.Split(words, " ") {
			decoded.WriteString(MORSE_CODE[letter])
		}
	}

	return decoded.String()
}
