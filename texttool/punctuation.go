package texttool

import (
	"regexp"
)

func Punctuation(s string) string {

	pat1 := regexp.MustCompile(`\s+([.,;:?!]+)`)

	s = pat1.ReplaceAllString(s, "$1")

	pat2 := regexp.MustCompile(`([;,.!?:]+)([^;.?:!\s])`)
	s = pat2.ReplaceAllString(s, "$1 $2")

	return s
}
