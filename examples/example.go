package main

import (
	"cmdline"
	"flag"
	"fmt"
	"strings"
)

const cmdName = "phone"

var commonNumberArg = cmdline.NewArg("number",
	"(123-456-7890) the phone number to call")

//// Call subcommand definitions ////
const cmdCallName = "call"

var (
	cmdCallFlags = flag.NewFlagSet(cmdCallName, flag.ExitOnError)
	speaker      = cmdCallFlags.Bool("speaker", false,
		"(bool) whether the speaker phone should be used")
)

var cmdCallArgs = []*cmdline.Arg{commonNumberArg}

type cmdCallRunner struct {
	calledHistory []string
}

func (c *cmdCallRunner) Run(args []string) error {
	// No need to check length of args here: the arg definition provided
	// guarantees it will at least have a length 1
	var speakerMsg string
	if *speaker {
		speakerMsg = " with speaker phone"
	}
	number := args[0]
	fmt.Printf("Dialing %s%s...\n", number, speakerMsg)
	c.calledHistory = append(c.calledHistory, number)
	return nil
}

//// Text subcommand definitions ////
const cmdTextName = "text"

var (
	cmdTextFlags = flag.NewFlagSet(cmdTextName, flag.ExitOnError)
	signature    = cmdTextFlags.String("signature", "",
		"(string) signature to use at the end of the message")
)

var cmdTextArgs = []*cmdline.Arg{
	commonNumberArg,
	cmdline.NewArg("message", "(string) the message to send"),
}

type cmdTextRunner struct{}

func (c *cmdTextRunner) Run(args []string) error {
	// No need to check length of args here: the arg definition provided
	// guarantees it will at least have a length 2
	number := args[0]
	message := strings.Join(args[1:], " ") // Just in case they forgot to put second arg in quotes...
	fmt.Printf("Texted %s this message: %s%s\n", number, message, *signature)
	return nil
}

func main() {
	callCmd := cmdline.NewCmd(cmdCallName, &cmdCallRunner{}, cmdCallArgs,
		cmdCallFlags)
	textCmd := cmdline.NewCmd(cmdTextName, &cmdTextRunner{}, cmdTextArgs,
		cmdTextFlags)

	mainCmd := cmdline.NewCmd(cmdName, nil, nil, nil)
	mainCmd.RegisterSub(callCmd)
	mainCmd.RegisterSub(textCmd)

	cmdline.Execute(mainCmd) // It's easy to forget this line ... but don't.
}
