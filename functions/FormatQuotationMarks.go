package goReloaded
import (
	"regexp"
)

func FormatQuotationMarks(text string) string {
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
