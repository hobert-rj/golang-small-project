package main

import (
	"fmt"
	"io/fs"
	"os"
)

type directory struct {
	path    string
	content []fs.DirEntry
}

func ReadDir(path string) directory {
	var res directory
	var err error
	res.path = path
	res.content, err = os.ReadDir(path)
	if err != nil {
		fmt.Println("err:\n", err)
	}
	return res
}
