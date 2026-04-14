package main

import (
	"fmt"
	"go-reloaded/texttool"
	"os"
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
	result = texttool.Command(result)

	result = texttool.Article(result)
	result = texttool.Base(result)

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file, ", err)
		return
	}
}
