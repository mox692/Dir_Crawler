package crawl

import (
	"flag"
	"fmt"
	"log"
	"os"
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

var (
	NO_FILE_FOUND        = "No file found."
	MULTIPLE_FILES_FOUND = "Multiple files found. Can't jump several dir."
)

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

	// modeによって処理を切り分け
	switch crawler.getMode() {
	case "list":
		err = crawler.WalkToList()
	case "jump":
		err = crawler.WalkToJump()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// 標準出力で描画
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

	// マッチするファイルが見つかったかの判定
	exist := c.isExist()
	if !exist {
		fmt.Println(NO_FILE_FOUND)
		return
	}

	switch c.mode {
	case "list":
		for _, v := range c.results {
			fmt.Printf("%s\n", v)
		}
	case "jump":
		if len(c.results) > 1 {
			fmt.Println(MULTIPLE_FILES_FOUND)
		}
		for _, v := range c.results {
			fmt.Printf("%s\n", v)
		}
	}
}
