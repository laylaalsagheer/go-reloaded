package goReloaded

import(
	"regexp"
	"strconv"
)
func ReplaceHexWithDecimal(text string) string {
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