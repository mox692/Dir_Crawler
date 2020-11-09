package main

import (
	"fmt"
)

func main() {
	fmt.Println("input your name.")

	var name string
	fmt.Scanf("%s", &name)

	if name == "motoyuki" {
		fmt.Println("hello!!")
	}
}
