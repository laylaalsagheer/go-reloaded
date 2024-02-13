package goReloaded
import (
	"strconv"
	"regexp"
	"strings"
	"fmt"
)

func ConvertToLower(text string) string {
	r := regexp.MustCompile(`\((low, ([0-9]+))\)`)                  //searching for the tag only
	m := regexp.MustCompile(`\w+|\(\w+,\s*\d+\)|[:!,\?.;'//\)\()]`) //\w+|\(\w+,\s*\d+\) //searching for the word or formated tag given in parentheses or any punctuation specified
	words := m.FindAllString(text, -1)                              // return all matched substrign, -1 for finding all the matches
	//loop over the array of matched string
	for index, word := range words {
		//ckeck if does the word contain the low tag with the number
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
			// num := int(chars[6] - '0') //change into number, potentional problem it works only with one digit
			if num > (len(words) - 1) {
				fmt.Println("invalid (low, <number>) format, your number is too big")
				break
			}
			for i := 1; i <= num; i++ {
				words[index-i] = strings.ToLower(words[index-i]) // apply changes based on the number that we got
			}
			words[index] = "" // remove the tag
		}
	}
	//join the elements based on the space as seperator, and get rid of the 2 spaces problem cause by line 149
	text = strings.ReplaceAll(strings.Join(words, " "), "  ", " ") //rewrite the line variable
	return text
}