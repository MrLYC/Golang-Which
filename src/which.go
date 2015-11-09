// which
package main

import (
	"fmt"
	"syscall"
)

func main() {
	argument := new(Argument)
	argument.ParseArgs()
	switch {
	case argument.option_help:
		{
			argument.PrintHelp()
			syscall.Exit(0)
		}
	case len(argument.filename) > 0:
		{
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
	default:
		{
			syscall.Exit(2)
		}
	}
}
