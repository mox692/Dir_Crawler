package main

import (
	"dirWalk/crawl"
	"os"
)

func main() {
	os.Exit(crawl.Run())
}
