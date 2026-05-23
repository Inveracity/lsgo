package main

import (
	"lsgo/internal"
	"os"
)

func main() {
	inputDir := os.Args[1:]

	dir := "."
	if len(inputDir) > 0 {
		dir = inputDir[0]
	}
	err := internal.ListDir(dir)
	if err != nil {
		panic(err)
	}
}
