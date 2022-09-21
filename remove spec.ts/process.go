package main

import (
	"fmt"
	"os"
	"strings"
)

type FilePaths []string

func (folder *directory) ProcessFolder() *FilePaths {
	var files FilePaths
	var process func(folder *directory)
	process = func(folder *directory) {
		for _, f := range folder.content {
			name := f.Name()
			path := folder.path + `\` + name
			if f.IsDir() {
				dir := ReadDir(path)
				process(&dir)
			} else {
				if strings.Contains(name, ".spec.ts") {
					files = append(files, path)
				}
			}
		}
	}
	process(folder)
	return &files
}

func (files *FilePaths) DeleteAll() *FilePaths {
	var success FilePaths
	for _, file := range *files {
		err := os.Remove(file)
		if err != nil {
			fmt.Println("err:\n", err)
		} else {
			success = append(success, file)
		}
	}
	return &success
}
