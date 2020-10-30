package main

import (
	"fmt"
	"os"
)

// cliでファイル名を入力

// 現在のdirでファイルを検索

// 一致が存在すれば 'yes' 存在しなければ 'no'
func main() {
	file := os.Args[1]
	fmt.Println(file)

}
