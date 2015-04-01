package main

import (
	"log"
	"regexp"
	"strings"
)

var validName *regexp.Regexp

func NameIsValid(name string) bool {
	if validName == nil {
		var err error
		validName, err = regexp.Compile(`(\w|(\w\.\w))+`)
		if err != nil {
			log.Panic(err)
		}
	}

	if validName.MatchString(name) {
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
