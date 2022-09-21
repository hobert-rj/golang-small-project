// Hossein Rajabi
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// for now only works for converting from px
const MinimumTolerance = 2
const DivideBy = 15.2
const target = "em"

var mainC chan string
var nonAlphanumericRegex *regexp.Regexp

func main() {
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	mainC = make(chan string)
	fmt.Println("Hossein Rajabi\nReading config...")
	var baseAddress, childAddress string
	{
		var inp string
		config := ReadFileAsString("config.txt")
		isKey := true
		var holdKey, holdVal string
		for _, char := range config {
			switch char {
			case '=':
				isKey = false
			case ';':
				switch strings.TrimSpace(holdKey) {
				case "app":
					inp = strings.TrimSpace(holdVal)
				case "base":
					baseAddress = strings.TrimSpace(holdVal)
				case "child":
					childAddress = strings.TrimSpace(holdVal)
				default:
					panic("Unknown property in config.")
				}
				isKey = true
				holdKey = ""
				holdVal = ""
			default:
				if isKey {
					holdKey += string(char)
				} else {
					holdVal += string(char)
				}
			}
		}
		if inp == "" {
			panic("app not defined")
		}
		if baseAddress == "" {
			panic("base not defined")
		}
		if childAddress == "" {
			panic("child not defined")
		}
		childAddress = inp + `\` + childAddress
		baseAddress = inp + `\` + baseAddress
	}
	filePaths := ReadFolder(childAddress)
	for key, fileGroup := range filePaths {
		go process(childAddress, key, fileGroup)
	}
	result := "export const AppIds = {\n"
	for i := 0; i < len(filePaths); i++ {
		key := <-mainC
		key = nonAlphanumericRegex.ReplaceAllString(key, "")
		val := strings.Title(key) + "Ids"
		result += fmt.Sprintf("  %s: %s,\n", key, val)
	}
	result += "}"
	err := os.WriteFile(baseAddress, []byte(result), 0666)
	if err != nil {
		fmt.Println("Error writing to file: ", baseAddress)
	}
	fmt.Println("Process finished.")
	fmt.Print("Enter to Exit...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()
}

func process(childAddress string, key string, pathGroup []string) {
	c := make(chan string)
	for _, file := range pathGroup {
		go processFile(file, &c)
	}
	result := "export const " + strings.Title(nonAlphanumericRegex.ReplaceAllString(key, "")) + "Ids = {\n"
	for i := 0; i < len(pathGroup); i++ {
		component := <-c
		component = nonAlphanumericRegex.ReplaceAllString(component, "")
		if component != "" {
			result += fmt.Sprintf("  %s: '%s',\n", component, component)
		}
	}
	result += "}"
	path := childAddress + `\` + key + `\` + key + ".id.ts"
	err := os.WriteFile(path, []byte(result), 0666)
	if err != nil {
		fmt.Println("Error writing to file: ", path)
	}
	mainC <- key
}

func processFile(path string, c *chan string) {
	fmt.Printf("Processing %s...\n", path)
	data := ReadFileAsString(path)
	var hold string
	componentFound := false
	classFound := false
	for _, char := range data {
		if unicode.IsSpace(char) {
			if hold != "" {
				if classFound {
					break
				} else if componentFound {
					if hold == "class" {
						classFound = true
					}
				} else if strings.Contains(hold, "@Component") {
					componentFound = true
				}
				hold = ""
			}
		} else {
			hold += string(char)
		}
	}
	*c <- hold
}
