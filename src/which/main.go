package main

import (
	"fmt"
	"syscall"
)

import (
	. "command"
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
			syscall.Exit(1)
		}
		syscall.Exit(0)
	}
	syscall.Exit(2)
}
