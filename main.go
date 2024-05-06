package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Use :: go run main.go input_file output_file")
		return
	}

	textBin, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	text := string(textBin)

	text = Grammar(text)
	text = Flags(text)
	text = Punctuation(text)
	text = Apostrophe(text)
	

	// Create a new file or truncate the existing file
	if strings.HasSuffix(os.Args[2], ".go") {
		log.Fatal("you cannot write in main ✋")
	}
	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
