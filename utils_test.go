package main

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, value1, value2 interface{}) {
	if !reflect.DeepEqual(value1, value2) {
		panic("AsserEqual failed: " + fmt.Sprint(value1) +
			" != " + fmt.Sprint(value2))
	}
}

func AssertUnequal(t *testing.T, value1, value2 interface{}) {
	if reflect.DeepEqual(value1, value2) {
		panic("AsserUnequal failed: " + fmt.Sprint(value1) +
			" != " + fmt.Sprint(value2))
	}
}

func TestNameIsValid(t *testing.T) {
	AssertEqual(t, true, NameIsValid("abc.dff.ere"))
	AssertEqual(t, true, NameIsValid("ert34"))
	AssertEqual(t, true, NameIsValid("Waw48.rtex.IOU74"))
	AssertEqual(t, true, NameIsValid(""))

	AssertEqual(t, false, NameIsValid(`.werwe`))
	AssertEqual(t, false, NameIsValid("retio*"))
	AssertEqual(t, false, NameIsValid("wev.poi%df.ete"))
	AssertEqual(t, false, NameIsValid("a.b  "))
	AssertEqual(t, false, NameIsValid("x.y z"))
	AssertEqual(t, false, NameIsValid("cvb.ewt."))
}

func TestNameToPath(t *testing.T) {
	AssertEqual(t, NameToPath("abc.eee"), "abc/eee")
	AssertEqual(t, NameToPath("er12"), "er12")
	AssertEqual(t, NameToPath("X.Y.Z"), "X/Y/Z")
}

func TestPathToName(t *testing.T) {
	AssertEqual(t, "abc.eee", PathToName("abc/eee"))
	AssertEqual(t, "er12", PathToName("er12"))
	AssertEqual(t, "X.Y.Z", PathToName("X/Y/Z"))
}
