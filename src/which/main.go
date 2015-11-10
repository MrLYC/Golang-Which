package main

import (
	"fmt"
	"syscall"
)

import (
	. "command"
)

const (
	OKStatus       = iota
	NotFoundStatus = iota
	ErrorStatus    = iota
)

func main() {
	argument := new(Argument)
	argument.ParseArgs()
	if len(argument.filename) > 0 {
		cmd := NewCommand(argument.filename)
		next := cmd.Find()
		existed := false
		for path := next(); len(path) > 0; path = next() {
			fmt.Println(path)
			existed = true
			if !argument.option_all {
				break
			}
		}
		if !existed {
			syscall.Exit(NotFoundStatus)
		}
		syscall.Exit(OKStatus)
	}
	syscall.Exit(ErrorStatus)
}
