package main

import (
	"goReloaded/functions"
	"fmt"
	"io/ioutil"
	"os"
)


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

func modifyText(inputFile, outputFile string) error {
	// Read input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Convert content to string
	text := string(content)

	//one tag convert
	modifiedText := goReloaded.OneTag(text)

	// Replace (hex) with decimal values
	modifiedText = goReloaded.ReplaceHexWithDecimal(modifiedText)

	// Replace (bin) with decimal values
	modifiedText = goReloaded.ReplaceBinWithDecimal(modifiedText)

	// Convert to capitalized
	modifiedText = goReloaded.ConvertToCapitalize(modifiedText)

	// Convert to uppercase
	modifiedText = goReloaded.ConvertToUpper(modifiedText)

	// Convert to lowercase
	modifiedText = goReloaded.ConvertToLower(modifiedText)

	// Format punctuations
	modifiedText = goReloaded.FormatPunctuation(modifiedText)

	// Format quotation marks
	modifiedText = goReloaded.FormatQuotationMarks(modifiedText)

	// Convert "a" to "an" based on the next word
	modifiedText = goReloaded.ConvertAtoAn(modifiedText)

	// Write modified text to output file
	err = ioutil.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		return err
	}

	return nil
}