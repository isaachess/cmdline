package main

import (
	"cmdline"
	"flag"
	"fmt"
)

const cmdCpName = "cp"

var cmdCpArgs = []*cmdline.Arg{
	cmdline.NewArg("sitename", "(string) the site you wish to copy the pw for"),
}

var cmdCpFlags *flag.FlagSet = nil

type cmdCpRunner struct {
	pws []*pw
}

func (c *cmdCpRunner) Run(args []string) error {
	var site = args[0]
	for _, pw := range c.pws {
		if pw.site == site {
			fmt.Println(pw.pw)
			return nil
		}
	}
	fmt.Println("No entry found for", site)
	return nil
}
