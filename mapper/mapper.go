package mapper

import (
	"fmt"
	"unicode"
)

// Interface describes the functions to modify a string
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// StringVal has the string type with position that needs to be transformed
type StringVal struct {
	Str string
	Pos int
}

//GetValueAsRuneSlice converts a string to slice of rune and returns it
func (v StringVal) GetValueAsRuneSlice() []rune {
	return []rune(v.Str)
}

// positionTracker keeps track of the position in a string
var positionTracker int = 1

// TransformRune transform a rune to upper/lowercase depending on the
// position of the string
func (v *StringVal) TransformRune(pos int) {
	str := v.GetValueAsRuneSlice()
	// upper case all the third character of the string
	if positionTracker == v.Pos {
		str[pos] = unicode.ToUpper(str[pos])
		positionTracker = 1
	} else {
		str[pos] = unicode.ToLower(str[pos])
		if unicode.IsLetter(str[pos]) {
			positionTracker++
		}
	}
	v.Str = string(str)
}

//MapString iterates through the interface and makes transformation.
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// String formats and returns string for string in StringVal type
func (v StringVal) String() string {
	return fmt.Sprint(v.Str)
}
