package main

import (
	"testing"
)

func TestCheckInterfaceConsistsOfMapsAndStrings(t *testing.T) {
	a := make(map[string]interface{})
	a["x"] = "abc"
	b := make(map[string]interface{})
	b["z"] = "fghfgh"
	b["t"] = "sdfewrwrwe"
	a["y"] = b
	AssertEqual(t, nil, CheckInterfaceConsistsOfMapsAndStrings(a))
	AssertEqual(t, nil, CheckInterfaceConsistsOfMapsAndStrings(b))
	AssertEqual(t, nil, CheckInterfaceConsistsOfMapsAndStrings("wads"))

	c := make(map[string]interface{})
	c["sdf"] = 5
	c["sdfs"] = "ewrewr"

	d := make(map[int]interface{})
	d[5] = "asadsad"

	e := make(map[string]interface{})
	e["wwrw"] = make([]int, 5)

	AssertEqual(t, false, CheckInterfaceConsistsOfMapsAndStrings(c) == nil)
	AssertEqual(t, false, CheckInterfaceConsistsOfMapsAndStrings(d) == nil)
	AssertEqual(t, false, CheckInterfaceConsistsOfMapsAndStrings(d) == nil)

}

func TestCreateGetInsertDelete(t *testing.T) {
	InitRootDirectory("/var/tmp/alex_config_tmp")
	DeleteNode("")
	defer DeleteNode("")

	tree := make(map[string]interface{})
	subtree1 := make(map[string]interface{})
	subtree2 := make(map[string]interface{})

	tree["i1"] = "v1"

	subtree1["i2"] = "v2"
	subtree1["i3"] = "v3"

	subtree2["i4"] = "v4"
	subtree2["i5"] = "v5"

	tree["subtree1"] = subtree1
	tree["subtree2"] = subtree2

	emptyTree, err := GetNode("", 0, 10000)
	AssertEqual(t, err, nil)
	AssertEqual(t, emptyTree, make(map[string]interface{}))
	AssertUnequal(t, tree, emptyTree)

	err = CreateNode("", tree)
	AssertEqual(t, err, nil)

	loadedTree, err := GetNode("", 0, 100000)
	AssertEqual(t, err, nil)
	AssertEqual(t, tree, loadedTree)

	err = CheckSubtreeMatchesValueStructure("", tree)
	AssertEqual(t, err, nil)

	loadedSubtree1, err := GetNode("subtree1", 0, 100000)
	AssertEqual(t, err, nil)
	AssertEqual(t, subtree1, loadedSubtree1)
	AssertUnequal(t, subtree2, loadedSubtree1)

	err = UpdateNode("subtree1/i3", "v3_2")
	AssertEqual(t, err, nil)

	loadedTree, err = GetNode("", 0, 10000)
	AssertEqual(t, err, nil)
	AssertUnequal(t, tree, loadedTree)

	tree["subtree1"].(map[string]interface{})["i3"] = "v3_2"
	AssertEqual(t, tree, loadedTree)

	changingTree := make(map[string]interface{})
	changingTree["subtree2"] = make(map[string]interface{})
	changingTree["subtree2"].(map[string]interface{})["i4"] = "v4_2"

	err = UpdateNode("", changingTree)
	AssertEqual(t, err, nil)

	loadedTree, err = GetNode("", 0, 10000)
	AssertEqual(t, err, nil)
	AssertUnequal(t, tree, loadedTree)

	tree["subtree2"].(map[string]interface{})["i4"] = "v4_2"
	AssertEqual(t, tree, loadedTree)

	err = DeleteNode("subtree1")
	AssertEqual(t, err, nil)

	err = CheckSubtreeMatchesValueStructure("", tree)
	AssertUnequal(t, err, nil)

	loadedTree, err = GetNode("", 0, 10000)
	AssertEqual(t, err, nil)
	AssertUnequal(t, tree, loadedTree)

	delete(tree, "subtree1")
	AssertEqual(t, tree, loadedTree)
}
