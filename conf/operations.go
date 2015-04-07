//Package conf implements operations for working with configurations
//represented by a hierarchical data structure.
//It stores nodes as files and use folders to describe the hierarchy.
//It allows to get and set entire directories using maps (map[string]interface{}).
//All the nodes values are strings.
package conf

import (
	"errors"
	"log"
)

//GetNode returns value of the node and all its subnodes.
//maxDepth defines the maximum recursion depth.
//Nodes values are string.
//Nodes hierarchy described via maps (map[string]interface{})
func GetNode(path string, maxDepth int) (value interface{}, err error) {

	if !FileOrDirExists(path) {
		return "", errors.New(ERROR_NODE_DOES_NOT_EXIST)
	}

	if IsDir(path) {
		subnodes := ReadDir(path)
		result := make(map[string]interface{})
		if maxDepth == 0 {
			return result, nil
		}
		var err1 error
		for _, node := range subnodes {
			result[node.Name()], err1 = GetNode(path+"/"+node.Name(), maxDepth-1)
			if err1 != nil {
				log.Panic(err1)
			}
		}
		return result, nil
	} else {
		return ReadFile(path), nil
	}

}

//CreateNode creates the node and all the subnodes described in the value parameter.
//Nodes values should be string.
//Nodes hierarchy should be described via maps (map[string]interface{})
//CreateNode returns an error if the specified node already exists.
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

//DeleteNode the specified node and all the subnodes.
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

//UpdateNode updates values of the specified node and all the subnodes.
//It updates only existing nodes.
//If the node or any or the subnodes described in value parameter
//doesn't exist, it returns an error.
//Nodes values should be string.
//Nodes hierarchy should be described via maps (map[string]interface{}).
func UpdateNode(path string, value interface{}) error {
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

//CheckSubtreeMatchesValueStructure checks if the specified value matches
//the tree structure.
//If the node or any or the subnodes described in value parameter
//doesn't exist, it returns an error.
//Nodes hierarchy should be described via maps (map[string]interface{}).
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

//CheckInterfaceConsistsOfMapsAndStrings checks if the value parameter
//consists of maps(map[string]interface{}) and strings.
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
