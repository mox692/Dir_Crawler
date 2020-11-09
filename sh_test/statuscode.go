package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(Run())
}

func Run() int {
	fmt.Println("run!")
	return 2
}
