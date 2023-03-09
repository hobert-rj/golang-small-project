// Hossein Rajabi
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// for now only works for converting from px
const MinimumTolerance = 2
const DivideBy = 15.2
const target = "em"

func main() {
	for {
		masterPath, err := ReadFolderPath()
		if err == nil && masterPath != "" {
			filePaths, err := ReadDeepFilePathWhen(masterPath, func(name string) bool {
				return strings.Contains(name, "Cheat Sheet.html")
			})
			if err == nil {
				m := make(map[int]string, len(filePaths))
				keys := make([]int, len(filePaths))
				ForEachFileAsString(filePaths, func(path string, file string) {
					s := strings.Split(path, `\`)
					name := s[len(s)-2]
					num_s := strings.Split(name, ".")[0]
					num, err := (strconv.ParseInt(num_s, 10, 32))
					if err != nil {
						fmt.Printf("Error converting to int: %s\n%v\n", num_s, err)
						return
					}
					file = fmt.Sprintf(`<h2>%s</h1>`, name) + file
					m[int(num)] = file
					keys = append(keys, int(num))
				})
				sort.Ints(keys)
				result := ""
				for _, key := range keys {
					result += m[key]
				}
				rPath := masterPath + `\merge.html`
				os.WriteFile(rPath, []byte(result), 0666)
				if err != nil {
					fmt.Println("Error writing to file: ", rPath)
				}
			}
			ExitOnEnter()
			break
		}
		fmt.Println("Try again.")
	}
}
