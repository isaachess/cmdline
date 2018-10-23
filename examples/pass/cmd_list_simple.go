package main

import (
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdListSimpleName = "simple"

var (
	cmdListSimpleArgs = args.NewArgSet()
	lsSite            = cmdListSimpleArgs.String("sitename",
		"(string) \"all\" or site name")
)

var cmdListSimpleFlags *flag.FlagSet

type cmdListSimpleRunner struct {
	pws []*pw
}

func (c *cmdListSimpleRunner) Run() error {
	printAll := *lsSite == "all"
	for _, pass := range c.pws {
		if printAll || pass.site == *lsSite {
			fmt.Println("Site", pass.site)
		}
	}
	return nil
}
