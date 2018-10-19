package main

import (
	"cmdline"
	"flag"
	"fmt"
)

const cmdListSimpleName = "simple"

var cmdListSimpleArgs = []*cmdline.Arg{
	cmdline.NewArg("sitename", "(string) \"all\" or site name"),
}

var cmdListSimpleFlags *flag.FlagSet

type cmdListSimpleRunner struct {
	pws []*pw
}

func (c *cmdListSimpleRunner) Run(args []string) error {
	site := args[0]
	printAll := site == "all"
	for _, pass := range c.pws {
		if printAll || pass.site == site {
			fmt.Println("Site", pass.site)
		}
	}
	return nil
}
