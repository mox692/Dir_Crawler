package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walk() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	fmt.Println(err)
}

func main() {
	walk()
}
