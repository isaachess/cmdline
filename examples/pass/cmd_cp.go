package main

import (
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdCpName = "cp"

var (
	cmdCpArgs = args.NewArgSet()
	cpSite    = cmdCpArgs.String("sitename",
		"(string) the site you wish to copy the pw for")
)

var cmdCpFlags *flag.FlagSet = nil

type cmdCpRunner struct {
	pws []*pw
}

func (c *cmdCpRunner) Run() error {
	for _, pw := range c.pws {
		if pw.site == *cpSite {
			fmt.Println(pw.pw)
			return nil
		}
	}
	fmt.Println("No entry found for", *cpSite)
	return nil
}
