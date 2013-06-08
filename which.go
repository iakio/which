package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path_env := os.Getenv("PATH")
	path_list := strings.Split(path_env, string(os.PathListSeparator))
	prognames := os.Args[1:]

	for _, progname := range prognames {
		for _, path := range path_list {
			fullpath := filepath.Join(path, progname)
			_, err := os.Stat(fullpath)
			if err == nil {
				fmt.Println(fullpath)
			}
		}
	}
}
