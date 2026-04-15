package texttool

import (
	"strings"
)

func Command(s string) string {

	text := strings.Fields(s)

	for r := 0; r < len(text); r++ {

		if r-1 > 0 {

			switch strings.ToLower(text[r]) {
			case "(up)":
				text[r-1] = strings.ToUpper(text[r-1])
				text = append(text[:r], text[r+1:]...)
				r--
			case "(low)":
				text[r-1] = strings.ToLower(text[r-1])
				text = append(text[:r], text[r+1:]...)
				r--
			case "(cap)":
				text[r-1] = strings.ToUpper(text[r-1][:1]) +strings.ToLower(text[r-1][1:])
				text = append(text[:r], text[r+1:]...)
				r--
			}
		}

	}
	return strings.Join(text, " ")
}
