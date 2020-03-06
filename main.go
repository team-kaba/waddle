package main

import (
	"./pkg/waddle"
	"os"
)

func main() {
	// cwdにはコマンド起動時のディレクトリが入っている
	cwd, _ := os.Getwd()
	os.Exit(int(waddle.Run(cwd, os.Args[1:]...)))
}
