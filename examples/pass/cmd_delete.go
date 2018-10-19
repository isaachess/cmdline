package main

import (
	"cmdline"
	"flag"
	"fmt"
)

const cmdDeleteName = "delete"

var cmdDeleteArgs = []*cmdline.Arg{
	cmdline.NewArg("sitename", "(string) the name of the site to delete"),
}

var cmdDeleteFlags = flag.NewFlagSet(cmdDeleteName, flag.ExitOnError)
var force = cmdDeleteFlags.Bool("force", false, "whether to force delete")

type cmdDeleteRunner struct {
	pws []*pw
}

func (c *cmdDeleteRunner) Run(args []string) error {
	site := args[0]
	if *force {
		fmt.Println("force not supported, but thanks for requesting it")
	}
	var newPws []*pw
	for _, pass := range c.pws {
		if pass.site == site {
			continue
		}
		newPws = append(newPws, pass)
	}
	c.pws = newPws
	return nil
}
