package goReloaded
import (
	"strings"
	"regexp"
)

func OneTag(text string) string{
	//up
	r := regexp.MustCompile(`([a-z|A-Z]+) \((up)\)`) //searching for a word and up tag
	// convert
	text = r.ReplaceAllStringFunc(text, func(s string) string { // when we found the output of our search we replace it
		match := r.FindStringSubmatch(s) //return array of strings
		return strings.ToUpper(match[1]) //return the first second element in upper format
	})
	//low
	r = regexp.MustCompile(`([a-z|A-Z]+) \((low)\)`) //searching for a word and low tag
	// convert
	text = r.ReplaceAllStringFunc(text, func(s string) string {
		match := r.FindStringSubmatch(s) //return array of strings
		return strings.ToLower(match[1]) //return the first second element in lower format
	})
	//cap
	r = regexp.MustCompile(`([a-z|A-Z]+) \((cap)\)`) //searching for a word and up tag
	// convert
	text = r.ReplaceAllStringFunc(text, func(s string) string {
		match := r.FindStringSubmatch(s) //return array of strings
		return strings.Title(strings.ToLower(match[1]))   //return the first second element in cap (title) format
	})
	return text
}