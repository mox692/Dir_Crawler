package crawl

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

// Crawler はディレクトリをクロールするためのstructです。
// userから引き受けたコマンド等を格納したり、見つかったfileの格納を行います。
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
	// NoFileFound はkeywordにマッチするfileが見つからなかった事を知らせます。
	NoFileFound = "! No file found."
	// MultipleFilesFound はjumpコマンドを選択した際に、keywordにマッチしたfileが複数存在した事を知らせます。
	MultipleFilesFound = "! Multiple files found. Can't jump several dir."
)

// Run はフラグの取得や、walk関数が返した結果をoutput関数に引き渡します。
func Run() int {

	var (
		list string
		get  string
		jump string
	)
	flag.StringVar(&list, "list", "", "`list` lists files that match keyword.")
	flag.StringVar(&jump, "jump", "", "`jump` jumps to the directory that match keyword.")
	flag.StringVar(&get, "get", "", "`get` gets the files that match keyword, and copy it to current directory.")

	flag.Parse()

	// list,jump,getが重複していたらerrを返します。
	if list != "" && jump != "" || jump != "" && get != "" || get != "" && list != "" {
		fmt.Printf("choose only 1 option from `list`, `get`, `jump`\n")
		return 1
	}

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

// WalkToList は配下のディレクトリをクロールして、キーワードマッチしたファイルをresultsに格納していきます。
func (c *Crawler) WalkToList() error {

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
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
				return err
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			if strings.Contains(string(buf), c.serchWord) {
				c.results = append(c.results, path)
			}
			file.Close()
		}
		return nil
	})
	return err
}

// WalkToJump は配下のディレクトリをクロールして、キーワードマッチしたファイルが存在するDirを取得します。
func (c *Crawler) WalkToJump() error {

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
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
				return err
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			if strings.Contains(string(buf), c.serchWord) {
				c.path[info.Name()] = path
				dirName := strings.Replace(path, info.Name(), "", 1)
				c.results = append(c.results, dirName)
			}
			file.Close()
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
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
	exist := c.isExist()
	if !exist {
		return xerrors.Errorf(NoFileFound)
	}

	switch c.mode {
	case "list":
		c.write()
	case "jump":
		if len(c.results) > 1 {
			fmt.Println(MultipleFilesFound)
			fmt.Println("-------------------------------------")
			c.write()
			fmt.Println("-------------------------------------")
			return xerrors.Errorf(MultipleFilesFound)
		}
		c.write()
	}
	return nil
}

func (c *Crawler) write() {
	switch c.mode {
	case "jump":
		if len(c.results) == 1 {
			fmt.Printf("%s\n", c.results[0])
			return
		}

		for _, value := range c.path {
			fmt.Printf("%s\n", value)
		}
	default:
		for _, v := range c.results {
			fmt.Printf("%s\n", v)
		}
	}
}
