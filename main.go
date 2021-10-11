package main

import (
	"aspiration/mapper"
	"fmt"
	"unicode"
)

func main() {
	//using CapitalizeEveryThirdAlphanumericChar define in this package to capitalize
	// every third alphanumeric char of a string
	fmt.Println("using CapitalizeEveryThirdAlphanumericChar function define in the main package to" +
		"capitalize every third alphanumeric char in a string")
	str := CapitalizeEveryThirdAlphanumericChar("Aspiration.com")
	fmt.Println(str)

	fmt.Println("\nusing mapper package to capitalize every third alphanumeric char in a string")
	// using mapper package to capitalize every third alphanumeric char
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}

// NewSkipString initializes the StringVal struct with the given parameters
func NewSkipString(pos int, s string) mapper.StringVal {
	return mapper.StringVal{
		Str: s,
		Pos: pos,
	}
}

//CapitalizeEveryThirdAlphanumericChar capitalized every third alphanumeric character
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	r := []rune(s) //convert to slice of rune
	tracker := 1
	for i := range r {
		// upper case all the third character of the string
		if tracker == 3 {
			if !unicode.IsLetter(r[i]) {
				continue
			}
			r[i] = unicode.ToUpper(r[i])
			tracker = 1
		} else {
			r[i] = unicode.ToLower(r[i])
			if unicode.IsLetter(r[i]) {
				tracker++
			}
		}

	}
	return string(r)

}
