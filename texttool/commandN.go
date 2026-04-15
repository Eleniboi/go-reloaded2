 package texttool

// import (
// 	"strconv"
// 	"strings"
// )

// func CommandN(s string) string {

// 	text := strings.Fields(s)

// 	for r := 0; r < len(text); r++ {

// 		if strings.ToLower(text[r]) == "(up," && strings.HasSuffix(text[r+1], ")") {

// 			clean := strings.TrimSuffix(text[r+1], ")")

// 			num, _ := strconv.Atoi(clean)

// 			text = append(text[:r], text[r+2:]...)
// 			r--

// 			count := r - num + 1

// 			if count < 0 {
// 				count = 0
// 			}

// 			for j := count; j <= r; j++ {
// 				text[j] = strings.ToUpper(text[j])
// 			}

// 		}
// 		if strings.HasPrefix(text[r],"(up,") && strings.HasSuffix(text[r], ")") {

// 			clean := strings.TrimPrefix(text[r], "(up,")
// 			clean = strings.TrimSuffix(clean, ")")

// 			num, _ := 
// 		}
// 	}
// }
