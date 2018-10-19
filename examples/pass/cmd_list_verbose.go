package main

import (
	"cmdline"
	"flag"
	"fmt"
)

const cmdListVerboseName = "verbose"

var cmdListVerboseArgs = []*cmdline.Arg{
	cmdline.NewArg("sitename", "(string) \"all\" or site name"),
}

var cmdListVerboseFlags *flag.FlagSet

type cmdListVerboseRunner struct {
	pws []*pw
}

func (c *cmdListVerboseRunner) Run(args []string) error {
	site := args[0]
	printAll := site == "all"
	for _, pass := range c.pws {
		if printAll || pass.site == site {
			fmt.Println("Site", pass.site)
			fmt.Println("Pass", pass.pw)
		}
	}
	return nil
}
