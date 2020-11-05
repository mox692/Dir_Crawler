package crawl

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
	rootDir   string
	results   []string
	mode      string
	list      string
	jump      string
	get       string
	path      map[string]string
	serchWord string
	skipWalk  bool
	result    string
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
			// .git系は飛ばす
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
				c.results = append(c.results, info.Name())
				if c.mode == "jump" {
					c.path[info.Name()] = path
				}
			}
			file.Close()
		}
		return nil
	})
	if err != nil {
		log.Fatal("err: %w", err)
	}

	if c.jump != "" {
		c.results = append(c.results, c.jumpToDir())
		return c.results
	}
	return []string{}
}

func (c *Crawler) jumpToDir() string {

	if len(c.results) > 1 {
		c.results = append(c.results, "More than 2 files were found. Can't jump to several dir.")
		return ""
	}
	if len(c.results) == 0 {
		c.results = append(c.results, "No file found.")
		return ""
	}

	c.skipWalk = true

	path := c.path[c.results[0]]
	if c.isCurrentDir(path) {
		fmt.Printf("`%s` is in current directory.\n", path)
		return ""
	}

	path = strings.Replace(path, c.results[0], "", 1)
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("err: %s", err)
		return ""
	}

	cwd2, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %s", err)
		return ""
	}
	return cwd2
}

func (c *Crawler) isCurrentDir(path string) bool {
	if strings.Contains(path, "/") {
		return false
	}
	return true
}

// Run はフラグの取得や、walk関数が返した結果をoutput関数に引き渡します。
func Run() {

	// フラグでfilenameの取得
	var (
		list string
		get  string
		jump string
	)
	flag.StringVar(&list, "list", "", "`list` lists files that match keyword.")
	flag.StringVar(&jump, "jump", "", "`jump` jumps to the directory that match keyword.")
	flag.StringVar(&get, "get", "", "`get` gets the files that match keyword, and copy it to current directory.")

	flag.Parse()

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

	root, err := os.Getwd()
	if err != nil {
		log.Fatal("err: %w", err)
	}

	crawler := &Crawler{
		rootDir:   root,
		mode:      mode, // listやjumpなどのmodeを格納します。
		list:      list, // 実際のvalueが入ります。
		jump:      jump,
		get:       get,
		serchWord: serchWord, // 上記のvalueを一時的に代入します。
		path:      map[string]string{},
		skipWalk:  false,
	}

	// results := crawler.walk()
	switch crawler.getMode() {
	case "list":
		err = crawler.WalkToList()

	case "jump":
		err = crawler.WalkToJump()
	}
	// *******************Todo: err処理
	if err != nil {
		log.Fatal("err: %w", err)
	}
	// 見つかったかの判定
	ok := crawler.isExist()
	if !ok {
		fmt.Printf("fileが見つかりません。")
		return
	}

	crawler.output()

}

func (c *Crawler) getMode() string {
	return c.mode
}

func (c *Crawler) isExist() bool {
	switch len(c.results) {
	case 0:
		return false
	default:
		return true
	}
}

func (c *Crawler) output() {
	for _, v := range c.results {
		fmt.Printf("%s\n", v)
	}
}
