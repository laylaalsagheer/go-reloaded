package goReloaded
import (
	"strconv"
	"regexp"
	"strings"
	"fmt"
)

func ConvertToCapitalize(text string) string {
	r := regexp.MustCompile(`\((cap, ([0-9]+))\)`)             //searching for the tag only
	m := regexp.MustCompile(`\w+|\(\w+,\s*\d+\)|[:!,\?.;'//\)\(]`) //\w+[:!,\?.;//]?|\(\w+,\s*\d+\)
	words := m.FindAllString(text, -1)                         // return all matched substrign, -1 for finding all the matches
	//loop over the array of matched string
	for index, word := range words {
		//appling the same logic as before
		if r.FindStringSubmatch(word) != nil {
			//fixing the multiple numbers issue
			chars := []rune(word) //if it does make it an array of rune
			nums := chars[6 : len(chars)-1]
			str := string(nums)
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			// num := int(chars[6] - '0')
			if num > (len(words) - 1) {
				fmt.Println("invalid (cap, <number>) format, your number is too big")
				break
			}
			for i := 1; i <= num; i++ {
				words[index-i] = strings.Title(words[index-i])
			}
			words[index] = "" // remove the tag
		}
	}
	text = strings.ReplaceAll(strings.Join(words, " "), "  ", " ")
	return text
}