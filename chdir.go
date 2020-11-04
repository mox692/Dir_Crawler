package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// supervisor などで起動した時は相対パスではなく絶対パスで指定する
	// output, err := exec.Command("./bin/test").Output()
	err := exec.Command("sh", "path.sh").Run()
	if err != nil {
		fmt.Printf("erreeee")
		fmt.Printf(err.Error())
	}
}
