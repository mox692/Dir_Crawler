package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

func walk(input string) string {
	var fileName string = ""
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if strings.Contains(info.Name(), ".git") {
				return filepath.SkipDir
			}
		}

		if !(info.IsDir()) {
			file, err := os.Open(path)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
			}
			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
			}

			// fmt.Println("テキストの中身：", string(buf))
			if strings.Contains(string(buf), input) {
				fileName = info.Name()
			}

			file.Close()
		}
		// fmt.Printf("visited file or dir: %v\n", path)
		return nil
	})
	if err != nil {
		log.Fatal("err: %w", err)
	}
	return fileName
}

func main() {
	input := os.Args[1]
	result := walk(input)

	if result == "" {
		fmt.Printf("fileが見つかりません。")
	} else {
		fmt.Printf("file: %s", result)
	}

}
