package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func ExampleWalk() {
	// 新しいdirの名前が返る
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	fmt.Printf("new dir name: %s\n", tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// IsDir => dir稼働かを判定
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
	// Output:
	// On Unix:
	// visited file or dir: "."
	// visited file or dir: "dir"
	// visited file or dir: "dir/to"
	// visited file or dir: "dir/to/walk"
	// skipping a dir without errors: skip
}

func main() {
	ExampleWalk()
}
