package crawl

import (
	"flag"
	"fmt"
	"os"

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

var (
	NO_FILE_FOUND        = "No file found."
	MULTIPLE_FILES_FOUND = "Multiple files found. Can't jump several dir."
)

// Run はフラグの取得や、walk関数が返した結果をoutput関数に引き渡します。
func Run() int {

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
		return 1
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
		fmt.Println(err)
		return 1
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
		return 1
	}

	// 標準出力で描画
	err = crawler.output()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
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

func (c *Crawler) output() error {

	// マッチするファイルが見つかったかの判定
	exist := c.isExist()
	if !exist {
		return xerrors.Errorf(NO_FILE_FOUND)
	}

	switch c.mode {
	case "list":
		c.write()
	case "jump":
		if len(c.results) > 1 {
			fmt.Println(MULTIPLE_FILES_FOUND)
			fmt.Println("-------------------------------------")
			c.write()
			return xerrors.Errorf(MULTIPLE_FILES_FOUND)
		}
		c.write()
	}
	return nil
}

func (c *Crawler) write() {
	switch c.mode {
	case "jump":
		for _, value := range c.path {
			fmt.Printf("%s\n", value)
		}

	default:
		for _, v := range c.results {
			fmt.Printf("%s\n", v)
		}
	}
}
