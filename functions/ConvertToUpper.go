package goReloaded
import (
	"strconv"
	"regexp"
	"strings"
	"fmt"
)

func ConvertToUpper(text string) string {
	r := regexp.MustCompile(`\((up, ([0-9]+))\)`)              //searching for the tag only
	m := regexp.MustCompile(`\w+|\(\w+,\s*\d+\)|[:!,\?.;'//\)\(]`) //refer to line 138
	words := m.FindAllString(text, -1)                         // return all matched substrign, -1 for finding all the matches
	//loop over the array of matched string
	for index, word := range words {
		if r.FindStringSubmatch(word) != nil {
			//appling the same logic as before
			//fixing the multiple numbers issue
			chars := []rune(word)
			nums := chars[5 : len(chars)-1]
			str := string(nums)
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			if num > (len(words) - 1) {
				fmt.Println("invalid (up, <number>) format, your number is too big")
				break
			}
			// num := int(chars[5] - '0')//the index changed becouse low has 3 characters and up has 2, therefore the number got shifted by one back
			for i := 1; i <= num; i++ {
				words[index-i] = strings.ToUpper(words[index-i])
			}
			words[index] = "" // remove the tag
		}
	}
	text = strings.ReplaceAll(strings.Join(words, " "), "  ", " ")
	return text
}
