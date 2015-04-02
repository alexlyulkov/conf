package main

import (
	"log"
	"regexp"
	"strings"
)

var validNameRe *regexp.Regexp

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

func NameToPath(name string) string {
	return strings.Replace(name, ".", "/", -1)
}

func PathToName(name string) string {
	return strings.Replace(name, "/", ".", -1)
}
