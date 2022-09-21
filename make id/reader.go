package main

import (
	"fmt"
	"os"
	"strings"
)

func AllowedFormats(format *string) bool {
	l := []string{"ts"}
	for _, f := range l {
		if f == *format {
			return true
		}
	}
	return false
}

func ReadFolder(inp string) map[string][]string {
	files := make(map[string][]string)
	fs, err := os.ReadDir(inp)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
	} else {
		for _, dir := range fs {
			if dir.IsDir() {
				name := dir.Name()
				files[name] = ReadFolderChild(inp + `\` + name)
			}
		}
	}
	return files
}

func ReadFolderChild(inp string) []string {
	var files []string
	fs, err := os.ReadDir(inp)
	if err != nil {
		fmt.Println("Error reading directory: ", inp)
		panic(err)
	} else {
		for _, dir := range fs {
			name := dir.Name()
			if dir.IsDir() {
				new := ReadFolderChild(inp + `\` + name)
				files = append(files, new...)
			} else {
				s := strings.Split(name, ".")
				if !AllowedFormats(&s[len(s)-1]) {
					continue
				}
				files = append(files, inp+`\`+name)
			}
		}
	}
	return files
}

func ReadFileAsString(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", path)
		panic(err)
	}
	return string(bytes)
}
