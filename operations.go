package main

import (
	"errors"
	"log"
)

//GetNode returns value of the node and its all subnotes
//subnotes are described using maps
func GetNode(path string, curDepth int, maxDepth int) (value interface{}, err error) {

	if !FileOrDirExists(path) {
		return "", errors.New(ERROR_NODE_DOES_NOT_EXIST)
	}

	if IsDir(path) {
		subnodes := ReadDir(path)
		result := make(map[string]interface{})
		if curDepth == maxDepth {
			return result, nil
		}
		var err1 error
		for _, node := range subnodes {
			result[node.Name()], err1 = GetNode(path+"/"+node.Name(),
				curDepth+1, maxDepth)
			if err1 != nil {
				log.Panic(err1)
			}
		}
		return result, nil
	} else {
		return ReadFile(path), nil
	}

}

//creates the node and all the subnotes described in value parameter
//subnodes should be described using maps
func CreateNode(path string, value interface{}) error {

	if len(path) != 0 && FileOrDirExists(path) {
		return errors.New(ERROR_NODE_ALREADY_EXISTS)
	}

	switch value := value.(type) {
	case string:
		WriteFile(path, string(value))
	case map[string]interface{}:
		if len(path) != 0 {
			MkDir(path)
		}
		for item, val := range value {
			err := CreateNode(path+"/"+item, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DeleteNode(path string) error {
	if len(path) == 0 {
		subnodes := ReadDir("")
		var err1 error
		for _, node := range subnodes {
			err1 = DeleteNode(path + "/" + node.Name())
			if err1 != nil {
				log.Panic(err1)
			}
		}
		return nil
	}
	if FileOrDirExists(path) {
		DeleteFileOrDir(path)
		return nil
	} else {
		return errors.New(ERROR_NODE_DOES_NOT_EXIST)
	}
}

func UpdateNode(path string, value interface{}) error {
	//updates nodes and all the subnotes described in value parameter
	//subnodes should be described using maps
	//all the nodes should exist

	if !FileOrDirExists(path) {
		return errors.New(ERROR_NODE_DOES_NOT_EXIST + " " + path)
	}

	switch value := value.(type) {
	case string:
		WriteFile(path, string(value))
	case map[string]interface{}:
		for item, val := range value {
			err := UpdateNode(path+"/"+item, val)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func CheckSubtreeMatchesValueStructure(path string,
	value interface{}) error {

	if !FileOrDirExists(path) {
		return errors.New(ERROR_NODE_DOES_NOT_EXIST + " " + PathToName(path))
	}

	switch value := value.(type) {
	case map[string]interface{}:
		if !IsDir(path) {
			return errors.New("Node " + PathToName(path) + " should have string value")
		}
		for item, val := range value {
			err := CheckSubtreeMatchesValueStructure(path+"/"+item, val)
			if err != nil {
				return err
			}
		}
	case string:
		if IsDir(path) {
			return errors.New("Node " + PathToName(path) + " can't have a string value, it is a subtree")
		}
	}

	return nil
}

func CheckInterfaceConsistsOfMapsAndStrings(value interface{}) error {

	switch value := value.(type) {
	case map[string]interface{}:
		for _, val := range value {
			err := CheckInterfaceConsistsOfMapsAndStrings(val)
			if err != nil {
				return err
			}
		}
		break
	case string:
		break
	default:
		return errors.New(ERROR_IVALID_VALUES_STRUCTURE)
	}

	return nil
}
