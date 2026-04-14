package texttool

import (
	"strings"
)

func Article(s string) string {

	text := strings.Fields(s)

	for r := 0; r < len(text); r++ {

		switch text[r] {
		case "a":
			if text[r] == "a" && strings.ContainsAny("aeiouhAEIOUH", text[r+1][:1]) {
				text[r] = "an"
			}
		case "A":
			if text[r] == "A" && strings.ContainsAny("aeiouhAEIOUH", text[r+1][:1]) {
				text[r] = "An"
			}
		case "an":
			if text[r] == "an" && !strings.ContainsAny("aeiouhAEIOUH", text[r+1][:1]) {
				text[r] = "a"
			}
		case "An":
			if text[r] == "An" && !strings.ContainsAny("aeiouhAEIOUH", text[r+1][:1]) {
				text[r] = "A"
			}

		}
	}
	return strings.Join(text, " ")
}
