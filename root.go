package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

type Crawler struct {
	fileNames []string
	mode      string
	list      string
	jump      string
	get       string
	path      map[string]string
	serchWord string
	skipWalk  bool
}

func (c *Crawler) setMode() {

}

func (c *Crawler) walk() []string {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if c.skipWalk {
			return nil
		}

		// dirに対する処理
		if info.IsDir() {
			if strings.Contains(info.Name(), ".git") {
				return filepath.SkipDir
			}
		}

		// fileに対する処理
		if !(info.IsDir()) {
			file, err := os.Open(path)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
			}
			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
			}

			if strings.Contains(string(buf), c.serchWord) {
				c.fileNames = append(c.fileNames, info.Name())

				if c.mode == "jump" {
					c.path[info.Name()] = path
				}
			}

			file.Close()
		}

		// それぞれのflag毎の処理
		// listの処理
		if c.list != "" {
			c.listFiles()
		}

		// jumpの処理
		if c.jump != "" {
			c.jumpToDir()
		}

		return nil
	})
	if err != nil {
		log.Fatal("err: %w", err)
	}
	return c.fileNames
}

func (c *Crawler) listFiles() {
	for _, v := range c.fileNames {
		fmt.Printf("fileName: %s\n", v)
	}
}
func (c *Crawler) jumpToDir() {

	if len(c.fileNames) > 1 {
		fmt.Printf("More than 2 files were found. Can't jump to several dir.\n")
		return
	}

	if len(c.fileNames) == 0 {
		return
	}
	path := c.path[c.fileNames[0]]

	path = strings.Replace(path, c.fileNames[0], "", 1)
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	c.skipWalk = true
	cwd2, err := os.Getwd()
	fmt.Println("now dir:", cwd2)
	return
}

func output(results []string) {
	for _, v := range results {
		fmt.Printf("%s\n", v)
	}
}

func main() {

	// フラグでfilenameの取得
	var (
		list string
		get  string
		jump string
	)
	flag.StringVar(&list, "list", "", "this is message!")
	flag.StringVar(&jump, "jump", "", "this is message!")
	flag.StringVar(&get, "get", "", "this is message!")

	flag.Parse()

	fmt.Printf("jump flag:%s\n", jump)
	fmt.Printf("list flag:%s\n", list)
	fmt.Printf("get  flag:%s\n", get)

	// list,jump,getが重複していたらerr
	if list != "" && jump != "" || jump != "" && get != "" || get != "" && list != "" {
		fmt.Printf("choose only 1 option from `list`, `get`, `jump`\n")
		return
	}

	// modeをセット
	var mode, serchWord string
	if list != "" {
		mode = "list"
		serchWord = list
	}
	if jump != "" {
		mode = "jump"
		serchWord = jump
	}
	if get != "" {
		mode = "get"
		serchWord = get
	}

	crawler := &Crawler{
		mode:      mode,
		list:      list,
		jump:      jump,
		get:       get,
		serchWord: serchWord,
		path:      map[string]string{},
		skipWalk:  false,
	}

	results := crawler.walk()
	if len(results) == 0 {
		fmt.Printf("fileが見つかりません。")
	} else {
		output(results)
	}

}
