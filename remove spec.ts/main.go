package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	unnecessaryWhiteSpaces := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	for {
		fmt.Println("/* Type \"End\" to end the program */\nEnter path of folder you want to process:")
		inp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err:\n", err)
			continue
		}
		inp = unnecessaryWhiteSpaces.ReplaceAllString(inp, "")
		if strings.ToLower(inp) == "end" {
			break
		}
		dir := ReadDir(inp)
		fileNames := dir.ProcessFolder()
		fmt.Println(`------files-----`)
		for _, name := range *fileNames {
			fmt.Println(name)
		}
		fmt.Println(`------files-----`)
		fmt.Println("Do you confirm deleting files listed above? type y to confirm")
		inp2, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err:\n", err)
			continue
		}
		inp2 = unnecessaryWhiteSpaces.ReplaceAllString(inp2, "")
		if inp2 == "y" {
			success := fileNames.DeleteAll()
			fmt.Println(`------Deleted files-----`)
			for _, name := range *success {
				fmt.Println(name)
			}
			fmt.Println(`------Deleted files-----`)
		}
	}
}
