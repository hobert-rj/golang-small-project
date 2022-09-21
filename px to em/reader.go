package main

import (
	"fmt"
	"os"
	"strings"
)

func AllowedFormats(format *string) bool {
	l := []string{"scss", "html"}
	for _, f := range l {
		if f == *format {
			return true
		}
	}
	return false
}

func ReadFolder(inp string) []string {
	var files []string
	fs, err := os.ReadDir(inp)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
	} else {
		for _, dir := range fs {
			name := dir.Name()
			if dir.IsDir() {
				new := ReadFolder(inp + `\` + name)
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
