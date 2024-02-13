package goReloaded
import (
	"strings"
)
func FormatPunctuation(text string) string {
	modifiedText := text

	// Replace ., ,, !, ?, :, and ; with proper spacing
	modifiedText = strings.ReplaceAll(modifiedText, " ,", ",")
	modifiedText = strings.ReplaceAll(modifiedText, " .", ".")
	modifiedText = strings.ReplaceAll(modifiedText, " !", "!")
	modifiedText = strings.ReplaceAll(modifiedText, " ?", "?")
	modifiedText = strings.ReplaceAll(modifiedText, " :", ":")
	modifiedText = strings.ReplaceAll(modifiedText, " ;", ";")
	modifiedText = strings.ReplaceAll(modifiedText, ":", ": ")
	// Format groups of punctuations
	modifiedText = strings.ReplaceAll(modifiedText, "...", "...")
	modifiedText = strings.ReplaceAll(modifiedText, "!!", "!!")
	modifiedText = strings.ReplaceAll(modifiedText, "??", "??")

	return modifiedText
}
