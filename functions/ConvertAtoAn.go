package goReloaded
import (
	"regexp"
)

func ConvertAtoAn(text string) string {
	modifiedText := text
	re := regexp.MustCompile(`\b(A|a)\s+(?i)([aeiou])`)
	modifiedText = re.ReplaceAllString(modifiedText, "An $2")

	return modifiedText
}

