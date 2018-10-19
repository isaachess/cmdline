package main

import (
	"cmdline"
	"flag"
	"fmt"
)

const cmdCreateName = "create"

var cmdCreateArgs = []*cmdline.Arg{
	cmdline.NewArg("sitename", "(string) the name of the site"),
	cmdline.NewArg("password", "(string) the password for the site"),
}

var cmdCreateFlags = flag.NewFlagSet(cmdCreateName, flag.ExitOnError)
var encrypt = cmdCreateFlags.Bool("encrypt", false, "whether to encrypt at rest")

type cmdCreateRunner struct {
	pws []*pw
}

func (c *cmdCreateRunner) Run(args []string) error {
	site := args[0]
	password := args[1]
	if *encrypt {
		fmt.Println("encrypt not supported, but thanks for requesting it")
	}
	c.pws = append(c.pws, &pw{site, password})
	return nil
}
