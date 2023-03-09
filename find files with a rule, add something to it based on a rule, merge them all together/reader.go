package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// func AllowedFormats(format string) bool {
// 	l := []string{"html"}
// 	for _, f := range l {
// 		if f == format {
// 			return true
// 		}
// 	}
// 	return false
// }

func ReadDeepFilePathWhen(inp string, fn func(name string) bool) ([]string, error) {
	var files []string
	fs, err := os.ReadDir(inp)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error reading directory: %s\n", inp), err)
		return files, err
	}
	for _, dir := range fs {
		name := dir.Name()
		if dir.IsDir() {
			new, err := ReadDeepFilePathWhen(inp+`\`+name, fn)
			if err == nil {
				files = append(files, new...)
			}
		} else {
			if !fn(name) {
				continue
			}
			files = append(files, inp+`\`+name)
		}
	}
	return files, nil
}

func ReadFolderPath() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hossein Rajabi\nEnter the folder address that you want to process...")
	inp, err := reader.ReadString('\n')
	if err == nil {
		inp = strings.TrimSpace(inp)
	}
	return inp, err
}

func ExitOnEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter to Exit...")
	reader.ReadLine()
}

func ForEachFileAsStringAsync(filePaths []string, fn func(file string, wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	for _, path := range filePaths {
		bytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file: ", path)
			continue
		}
		fmt.Printf("Processing %s...\n", path)
		file := string(bytes)
		wg.Add(1)
		go fn(file, &wg)
	}
	wg.Wait()
	fmt.Println("Process finished.")
}

func ForEachFileAsString(filePaths []string, fn func(path string, file string)) {
	for _, path := range filePaths {
		bytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file: ", path)
			continue
		}
		fmt.Printf("Processing %s...\n", path)
		file := string(bytes)
		fn(path, file)
	}
	fmt.Println("Process finished.")
}
