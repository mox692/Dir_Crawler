package crawl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

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
				return xerrors.Errorf("err: %w", err)
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
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

	_, err = c.isSpecific()
	if err != nil {
		return nil
	}

	return nil
}

func (c *Crawler) isSpecific() (bool, error) {
	if len(c.results) == 0 {
		return false, xerrors.New("no file found")
	}
	if len(c.results) > 1 {
		return false, xerrors.New("ファイルがたくさんあります。")
	}
	return true, nil
}
