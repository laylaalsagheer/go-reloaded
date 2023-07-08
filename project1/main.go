package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func replaceHexWithDecimal(text string) string {
	reHex := regexp.MustCompile(`([0-9A-Fa-f]+) \(hex\)`)
	modifiedText := reHex.ReplaceAllStringFunc(text, func(match string) string {
		hex := reHex.FindStringSubmatch(match)[1]
		dec, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			return match // Invalid hexadecimal, return the original match
		}
		return strconv.FormatInt(dec, 10)
	})
	return modifiedText
}

func replaceBinWithDecimal(text string) string {
	reBin := regexp.MustCompile(`([01]+) \(bin\)`)
	modifiedText := reBin.ReplaceAllStringFunc(text, func(match string) string {
		bin := reBin.FindStringSubmatch(match)[1]
		dec, err := strconv.ParseInt(bin, 2, 64)
		if err != nil {
			return match // Invalid binary, return the original match
		}
		return strconv.FormatInt(dec, 10)
	})
	return modifiedText
}

func convertToUpper(text string) string {
	reUp := regexp.MustCompile(`(\b\w+)\s\(up(\s*,\s*(\d+))?\)`)
	modifiedText := reUp.ReplaceAllStringFunc(text, func(match string) string {
		groups := reUp.FindStringSubmatch(match)
		word := groups[1]
		countStr := groups[3]
		if countStr == "" {
			return strings.ToUpper(word)
		}
		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			return match // Invalid count, return the original match
		}
		upperWord := strings.ToUpper(word)
		return strings.Repeat(upperWord, count)
	})
	return modifiedText
}

func convertToLower(text string) string {
	reLow := regexp.MustCompile(`(\b\w+)\s\(low(\s*,\s*(\d+))?\)`)
	modifiedText := reLow.ReplaceAllStringFunc(text, func(match string) string {
		groups := reLow.FindStringSubmatch(match)
		word := groups[1]
		countStr := groups[3]
		if countStr == "" {
			return strings.ToLower(word)
		}
		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			return match // Invalid count, return the original match
		}
		lowerWord := strings.ToLower(word)
		return strings.Repeat(lowerWord, count)
	})
	return modifiedText
}

func convertToCapitalize(text string) string {
	reCap := regexp.MustCompile(`(\b\w+)\s\(cap(\s*,\s*(\d+))?\)`)
	modifiedText := reCap.ReplaceAllStringFunc(text, func(match string) string {
		groups := reCap.FindStringSubmatch(match)
		word := groups[1]
		countStr := groups[3]
		if countStr == "" {
			return strings.Title(word)
		}
		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			return match // Invalid count, return the original match
		}
		capWord := strings.Title(word)
		return strings.Repeat(capWord, count)
	})
	return modifiedText
}

func formatPunctuation(text string) string {
	modifiedText := text

	// Replace ., ,, !, ?, :, and ; with proper spacing
	modifiedText = strings.ReplaceAll(modifiedText, " ,", ",")
	modifiedText = strings.ReplaceAll(modifiedText, " .", ".")
	modifiedText = strings.ReplaceAll(modifiedText, " !", "!")
	modifiedText = strings.ReplaceAll(modifiedText, " ?", "?")
	modifiedText = strings.ReplaceAll(modifiedText, " :", ":")
	modifiedText = strings.ReplaceAll(modifiedText, " ;", ";")

	// Format groups of punctuations
	modifiedText = strings.ReplaceAll(modifiedText, "...", "...")
	modifiedText = strings.ReplaceAll(modifiedText, "!!", "!!")
	modifiedText = strings.ReplaceAll(modifiedText, "??", "??")

	return modifiedText
}

func formatQuotationMarks(text string) string {
	modifiedText := text

	// Replace single quotation marks with proper spacing
	modifiedText = regexp.MustCompile(`\s'\s`).ReplaceAllString(modifiedText, " '")
	modifiedText = regexp.MustCompile(`\s'(\w+)`).ReplaceAllString(modifiedText, " '$1")
	modifiedText = regexp.MustCompile(`(\b\w+)'\s`).ReplaceAllString(modifiedText, "$1' ")

	// Remove spaces around single quotation marks
	modifiedText = regexp.MustCompile(`\s'`).ReplaceAllString(modifiedText, "'")
	modifiedText = regexp.MustCompile(`'\s`).ReplaceAllString(modifiedText, "'")

	// Replace multiple wordsbetween quotation marks with the same number of quotation marks
	modifiedText = regexp.MustCompile(`'\s+([^']+?)\s+'`).ReplaceAllString(modifiedText, "'$1'")

	return modifiedText
}

func convertAtoAn(text string) string {
	modifiedText := text
	re := regexp.MustCompile(`\b(A|a)\s+(?i)([aeiou])`)
	modifiedText = re.ReplaceAllString(modifiedText, "An $2")

	return modifiedText
}

func modifyText(inputFile, outputFile string) error {
	// Read input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Convert content to string
	text := string(content)

	// Replace (hex) with decimal values
	modifiedText := replaceHexWithDecimal(text)

	// Replace (bin) with decimal values
	modifiedText = replaceBinWithDecimal(modifiedText)

	// Convert to uppercase
	modifiedText = convertToUpper(modifiedText)

	// Convert to lowercase
	modifiedText = convertToLower(modifiedText)

	// Convert to capitalized
	modifiedText = convertToCapitalize(modifiedText)

	// Format punctuations
	modifiedText = formatPunctuation(modifiedText)

	// Format quotation marks
	modifiedText = formatQuotationMarks(modifiedText)

	// Convert "a" to "an" based on the next word
	modifiedText = convertAtoAn(modifiedText)

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
