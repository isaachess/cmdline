package main

import (
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdListVerboseName = "verbose"

var (
	cmdListVerboseArgs = args.NewArgSet()
	lvSite             = cmdListVerboseArgs.String("sitename",
		"(string) \"all\" or site name")
)

var cmdListVerboseFlags *flag.FlagSet

type cmdListVerboseRunner struct {
	pws []*pw
}

func (c *cmdListVerboseRunner) Run() error {
	printAll := *lvSite == "all"
	for _, pass := range c.pws {
		if printAll || pass.site == *lvSite {
			fmt.Println("Site", pass.site)
			fmt.Println("Pass", pass.pw)
		}
	}
	return nil
}
