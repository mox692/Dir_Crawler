package crawl

import (
	"fmt"
	"os"
	"testing"
)

func Test_WalkToList(t *testing.T) {

	// example struct作成
	crawlers := []*Crawler{
		{
			mode:      "list",
			serchWord: "testttttt",
		},
		{
			mode:      "list",
			serchWord: "4324ffffff",
		},
		{
			mode:      "list",
			serchWord: "fdsafas",
		},
		{
			mode:      "list",
			serchWord: "iiii",
		},
	}

	file1, err := os.Create("test1.txt")
	if err != nil {
		t.Errorf("%v", err)
	}
	file2, err := os.Create("test2.txt")
	if err != nil {
		t.Errorf("%v", err)
	}
	file3, err := os.Create("test3.txt")
	if err != nil {
		t.Errorf("%v", err)
	}
	file4, err := os.Create("test4.txt")
	if err != nil {
		t.Errorf("%v", err)
	}

	if err := os.MkdirAll("tmp/hoge/huga", 0777); err != nil {
		t.Errorf("%v", err)
	}

	err = os.Rename("test1.txt", "tmp/test1.txt")
	err = os.Rename("test2.txt", "tmp/hoge/test2.txt")
	err = os.Rename("test3.txt", "tmp/hoge/test3.txt")
	err = os.Rename("test4.txt", "tmp/hoge/huga/test4.txt")
	if err != nil { // 見やすさのため、1つのerrorで代表させています。
		t.Errorf("%v", err)
	}

	_, err = file1.Write([]byte("testttttt"))
	_, err = file2.Write([]byte("4324ffffff"))
	_, err = file3.Write([]byte("fdsafas"))
	_, err = file4.Write([]byte("iiii"))
	if err != nil { // 見やすさのため、1つのerrorで代表させています。
		t.Errorf("%v", err)
	}

	expect := []string{
		"test1.txt",
		"hoge/test2.txt",
		"hoge/test3.txt",
		"hoge/huga/test4.txt",
	}

	err = os.Chdir("tmp")
	if err != nil {
		t.Errorf("%v", err)
	}

	for i, c := range crawlers {

		err := c.WalkToList()

		if err != nil {
			t.Errorf("%v", err)
		}

		if len(c.results) != 1 {
			t.Errorf("len(c.results) is expected %d, got %d (%+v)", 1, len(c.results), c.results)
		}

		if c.results[0] != expect[i] {
			t.Errorf("c.results is expected %s, got %s", expect[i], c.results[0])
		}
	}

	// リソース削除
	err = os.Chdir("..")
	if err := os.RemoveAll("tmp"); err != nil {
		fmt.Println(err)
	}
}
