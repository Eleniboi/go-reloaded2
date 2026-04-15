package main

import (
	"fmt"
	"go-reloaded/texttool"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Invalid input: usage go run . input.txt output.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	file, err := os.ReadFile(inputfile)

	if err != nil {
		fmt.Println("Error reading file, ", err)
		return
	}

	result := string(file)
	result = processor(result)

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file, ", err)
		return
	}
}

func processor(s string) string {

	text := strings.Split(s, "\n")

	for i := 0; i < len(text); i++ {

		text[i] = texttool.Command(text[i])
		text[i] = texttool.Base(text[i])
		//text[i] = texttool.CommandN(text[i])
		text[i] = texttool.Article(text[i])
		text[i] = texttool.Punctuation(text[i])
		text[i] = texttool.Quote(text[i])

	}
	return strings.Join(text, "\n")
}
