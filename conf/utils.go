package conf

import (
	"log"
	"regexp"
	"strings"
)

var validNameRe *regexp.Regexp

//NameIsValid returns true if the specified string is a valid node name.
//Valid node name should consist only of english literals and numbers.
//Node and its parents names should be separated by dots.
func NameIsValid(name string) bool {
	if validNameRe == nil {
		var err error
		validNameRe, err = regexp.Compile(`^(\w|(\w\.\w))*$`)
		if err != nil {
			log.Panic(err)
		}
	}

	if validNameRe.MatchString(name) {
		return true
	} else {
		return false
	}
}

//NameToPath converts a valid dot-separated node name to
//the corresponding file path
func NameToPath(name string) string {
	return strings.Replace(name, ".", "/", -1)
}

//PathToName converts a file path to the corresponding dot-separated
//node name
func PathToName(name string) string {
	return strings.Replace(name, "/", ".", -1)
}
