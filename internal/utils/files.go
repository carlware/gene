package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func FileToString(filename string) (string, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(buf), err
}

func CreateFolders(filename string) error {
	paths := strings.Split(filename, "/")
	path := strings.Join(paths[0:len(paths)-1], "/")
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

func BytesToFile(filename string, content []byte) error {
	panic("unimplemented")
}
