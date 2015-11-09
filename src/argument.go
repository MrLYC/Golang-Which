// argument
package main

import (
	"flag"
	"fmt"
)

type Argument struct {
	filename    string
	option_all  bool
	option_help bool
}

func (this *Argument) ParseArgs() {
	filename := flag.String("filename", "", "")
	option_all := flag.Bool(
		"a", false, "print all matching pathnames of each argument",
	)
	option_s_help := flag.Bool(
		"h", false, "show this help message and exit",
	)
	option_help := flag.Bool(
		"help", false, "show this help message and exit",
	)
	flag.Parse()

	this.filename = *filename
	this.option_all = *option_all
	this.option_help = *option_s_help || *option_help
}

func (this *Argument) PrintHelp() {
	fmt.Println("usage: which - locate a command [-h] [-a] filename\n")
	fmt.Println("positional arguments:")
	fmt.Println("  filename")
	fmt.Println("\n")
	fmt.Println("optional arguments:")
	fmt.Println("  -a          print all matching pathnames of each argument")
	fmt.Println("  -h, --help  show this help message and exit")
}
