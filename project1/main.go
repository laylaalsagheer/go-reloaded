package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func modifyText(inputFile, outputFile string) error {
	// Read input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Convert content to string
	text := string(content)

	// Replace (hex) with decimal values
	re := regexp.MustCompile(`([0-9A-Fa-f]+) \(hex\)`)
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		hex := re.FindStringSubmatch(match)[1]
		dec, _ := strconv.ParseInt(hex, 16, 64)
		return strconv.FormatInt(dec, 10)
	})

	// Write modified text to output file
	err = ioutil.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Check if input and output file names are provided as command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-file> <output-file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := modifyText(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println("Text modified successfully.")
}
