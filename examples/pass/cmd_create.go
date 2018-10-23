package main

import (
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdCreateName = "create"

var (
	cmdCreateArgs = args.NewArgSet()
	createSite    = cmdCreateArgs.String("sitename",
		"(string) the name of the site")
	password = cmdCreateArgs.String("password",
		"(string) the password for the site")
)

var cmdCreateFlags = flag.NewFlagSet(cmdCreateName, flag.ExitOnError)
var encrypt = cmdCreateFlags.Bool("encrypt", false,
	"whether to encrypt at rest")

type cmdCreateRunner struct {
	pws []*pw
}

func (c *cmdCreateRunner) Run() error {
	if *encrypt {
		fmt.Println("encrypt not supported, but thanks for requesting it")
	}
	c.pws = append(c.pws, &pw{*createSite, *password})
	return nil
}
