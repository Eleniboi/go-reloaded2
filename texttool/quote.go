package texttool

import (
	"strings"
)

func Quote(s string) string {

	result := strings.Split(s, "'")

	for i := 0; i < len(result); i++ {

		if i%2 == 1 {

			result[i] = strings.TrimSpace(result[i])
		}
	}
	return strings.Join(result, "'")
}
