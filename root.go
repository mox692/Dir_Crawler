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

func walk(input string) []string {
	var fileNames []string
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
				fileNames = append(fileNames, info.Name())
			}

			file.Close()
		}
		// fmt.Printf("visited file or dir: %v\n", path)
		return nil
	})
	if err != nil {
		log.Fatal("err: %w", err)
	}
	return fileNames
}

func output(results []string) {
	for _, v := range results {
		fmt.Printf("%s\n", v)
	}
}

type RunMode string

type Crawler struct {
	mode RunMode
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Please input option :(")
		return
	}

	if len(os.Args) >= 2 {
		fmt.Println("too much option given :(")
		return
	}

	input := os.Args[1] // `jump` or `get`

	switch os.Args[1] {
	case "jump":

	}
	results := walk(input)
	if len(results) == 0 {
		fmt.Printf("fileが見つかりません。")
	} else {
		output(results)
	}

}
