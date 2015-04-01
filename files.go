package main

import (
	"io/ioutil"
	"log"
	"os"
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
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	log.Panic(err)
}

func ReadFile(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return string(buf)
}

func WriteFile(path string, value string) {
	err := ioutil.WriteFile(path, ([]byte)(path), 0755)
	if err != nil {
		log.Panic(err)
	}
}

func ReadDir(path string) []os.FileInfo {
	entries, err := ioutil.ReadDir(path)
	return entries
}
