package conf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//RootDir is the working directory
var RootDir string

//InitRootDirectory Initializes and sets the working directory.
//All the dirs and files will be in that directory
func InitRootDirectory(rootDir string) {
	RootDir = rootDir

	if err := os.MkdirAll(RootDir, os.ModeDir|0755); err != nil {
		log.Panic(err)
	}
	if err := os.Chdir(RootDir); err != nil {
		log.Panic(err)
	}
}

//FileOrDirExists returns true if the specified file or dir exists
func FileOrDirExists(path string) bool {
	path = "./" + path
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

	log.Panic(err)
	return false
}

//IsDir returns true if the object on the specified path is a directory
func IsDir(path string) bool {
	path = "./" + path
	fileInfo, err := os.Stat(path)

	if err != nil {
		log.Panic(err)
	}

	return fileInfo.IsDir()
}

//ReadFile returns the content of the file as a string
func ReadFile(path string) string {
	path = "./" + path
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return string(buf)
}

//WriteFile erases the specified file and writes the value string to it
func WriteFile(path string, value string) {
	path = "./" + path
	dir := filepath.Dir(path)
	if !FileOrDirExists(dir) {
		MkDir(dir)
	}
	err := ioutil.WriteFile(path, ([]byte)(value), 0755)
	if err != nil {
		log.Panic(err)
	}
}

//ReadDir returns FileInfo description of each
//file and directory in the specified directory
func ReadDir(path string) []os.FileInfo {
	path = "./" + path
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}
	return entries
}

//MkDir Creates a directory named path, along with any necessary parents
func MkDir(path string) {
	path = "./" + path
	err := os.MkdirAll(path, os.ModeDir|0755)
	if err != nil {
		log.Panic(err)
	}
}

//DeleteFileOrDir deletes the named file or directory
func DeleteFileOrDir(path string) {
	path = "./" + path
	if IsDir(path) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Panic(err)
		}
	} else {
		err := os.Remove(path)
		if err != nil {
			log.Panic(err)
		}
	}
}
