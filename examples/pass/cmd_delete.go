package main

import (
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdDeleteName = "delete"

var (
	cmdDeleteArgs = args.NewArgSet()
	deleteSite    = cmdDeleteArgs.String("sitename",
		"(string) the name of the site to delete")
)

var cmdDeleteFlags = flag.NewFlagSet(cmdDeleteName, flag.ExitOnError)
var force = cmdDeleteFlags.Bool("force", false, "whether to force delete")

type cmdDeleteRunner struct {
	pws []*pw
}

func (c *cmdDeleteRunner) Run() error {
	if *force {
		fmt.Println("force not supported, but thanks for requesting it")
	}
	var newPws []*pw
	for _, pass := range c.pws {
		if pass.site == *deleteSite {
			continue
		}
		newPws = append(newPws, pass)
	}
	c.pws = newPws
	return nil
}
