// argument
package main

import (
	"flag"
)

type Argument struct {
	filename   string
	option_all bool
}

func (this *Argument) ParseArgs() {
	flag.BoolVar(
		&(this.option_all), "a", false, "print all matching pathnames of each argument",
	)
	flag.Parse()
	this.filename = flag.Arg(0)
}
