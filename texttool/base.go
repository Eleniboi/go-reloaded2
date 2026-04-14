package texttool

import (
	"fmt"
	"strconv"
	"strings"
)

func Base(s string) string {

	text := strings.Fields(s)

	for r := 0; r < len(text); r++ {

		switch strings.ToLower(text[r]) {
		case "(hex)":
			n, err := strconv.ParseInt(text[r-1], 16, 64)
			if err != nil {
				continue
			}
			text[r-1] = fmt.Sprint(n)
			text = append(text[:r], text[r+1:]...)
			r--
		case "(bin)":
			n, err := strconv.ParseInt(text[r-1], 2, 64)
			if err != nil {
				continue
			}

			text[r-1] = fmt.Sprint(n)
			text = append(text[:r], text[r+1:]...)
			r--
		}

	}
	return strings.Join(text, " ")
}
