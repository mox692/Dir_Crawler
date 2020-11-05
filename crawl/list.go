package crawl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

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
				return xerrors.Errorf("err: %w", err)
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return xerrors.Errorf("err: %w", err)
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
