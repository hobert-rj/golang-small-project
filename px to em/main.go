// Hossein Rajabi
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

// for now only works for converting from px
const MinimumTolerance = 2
const DivideBy = 15.2
const target = "em"

var wg sync.WaitGroup

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Hossein Rajabi\nEnter folder address that you want to process...")
		inp, err := reader.ReadString('\n')
		inp = strings.TrimSpace(inp)
		if err == nil && inp != "" {
			filePaths := ReadFolder(inp)
			for _, file := range filePaths {
				wg.Add(1)
				go process(file)
			}
			wg.Wait()
			fmt.Println("Process finished.")
			fmt.Print("Enter to Exit...")
			reader.ReadLine()
			break
		}
		fmt.Println("Try again.")
	}
}

func process(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", path)
		wg.Done()
		return
	}
	fmt.Printf("Processing %s...\n", path)
	data := string(bytes)
	var result string
	var holdDigit string
	var holdStart int
	var holdEnd int
	for i, char := range data {
		if (unicode.IsDigit(char) || char == '.') && !(char == '0' && holdDigit == "") {
			if holdDigit == "" {
				holdStart = i
			}
			holdDigit += string(char)
		} else if holdDigit != "" {
			if char == 'p' && (data)[i+1] == 'x' {
				num, err := strconv.ParseFloat(holdDigit, 64)
				if err != nil {
					fmt.Printf("Error converting %s to float. %e", holdDigit, err)
				}
				if num > MinimumTolerance {
					r := num / DivideBy
					result += data[holdEnd:holdStart] + fmt.Sprintf("%.2f", r) + target
					holdEnd = i + 2
				}
				holdStart = 0
			}
			holdDigit = ""
		}
	}
	result += data[holdEnd:]
	os.WriteFile(path, []byte(result), 0666)
	if err != nil {
		fmt.Println("Error writing to file: ", path)
	}
	wg.Done()
}
