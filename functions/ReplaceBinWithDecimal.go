package goReloaded
import (
	"strconv"
	"regexp"
)
func ReplaceBinWithDecimal(text string) string {
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

