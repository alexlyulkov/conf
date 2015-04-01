package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var RootDir string

func InitRootDirectory(rootDir string) {
	RootDir = rootDir

	if err := os.MkdirAll(RootDir, os.ModeDir|0755); err != nil {
		log.Panic(err)
	}
	if err := os.Chdir(RootDir); err != nil {
		log.Panic(err)
	}
}

func DirExists(path string) bool {
	err := os.Chdir(path)
	defer os.Chdir(RootDir)

	if err == nil {
		return true
	}

	if !os.IsNotExist(err) {
		return false
	}

	log.Panic(err)
	return false
}

func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

	log.Panic(err)
	return false
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		log.Panic(err)
	}

	return fileInfo.IsDir()
}

func ReadFile(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return string(buf)
}

func WriteFile(path string, value string) {
	dir := filepath.Dir(path)
	if !FileOrDirExists(dir) {
		MkDir(dir)
	}
	err := ioutil.WriteFile(path, ([]byte)(value), 0755)
	if err != nil {
		log.Panic(err)
	}
}

func ReadDir(path string) []os.FileInfo {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}
	return entries
}

func MkDir(path string) {
	err := os.MkdirAll(path, os.ModeDir|0755)
	if err != nil {
		log.Panic(err)
	}
}

func DeleteFileOrDir(path string) {
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
