package main

import (
	"cmdline"
	"cmdline/args"
	"flag"
	"fmt"
)

const cmdName = "phone"

//// Call subcommand definitions ////
const cmdCallName = "call"

var (
	cmdCallFlags = flag.NewFlagSet(cmdCallName, flag.ExitOnError)
	speaker      = cmdCallFlags.Bool("speaker", false,
		"(bool) whether the speaker phone should be used")
)

var (
	cmdCallArgs = args.NewArgSet()
	callNumber  = cmdCallArgs.String("number",
		"(123-456-7890) the phone number to call")
)

type cmdCallRunner struct {
	calledHistory []string
}

func (c *cmdCallRunner) Run() error {
	var speakerMsg string
	if *speaker {
		speakerMsg = " with speaker phone"
	}
	fmt.Printf("Dialing %s%s...\n", *callNumber, speakerMsg)
	c.calledHistory = append(c.calledHistory, *callNumber)
	return nil
}

//// Text subcommand definitions ////
const cmdTextName = "text"

var (
	cmdTextFlags = flag.NewFlagSet(cmdTextName, flag.ExitOnError)
	signature    = cmdTextFlags.String("signature", "",
		"(string) signature to use at the end of the message")
)

var (
	cmdTextArgs = args.NewArgSet()
	textNumber  = cmdTextArgs.String("number",
		"(123-456-7890) the phone number to call")
	message = cmdTextArgs.String("message", "(string) the message to send")
)

type cmdTextRunner struct{}

func (c *cmdTextRunner) Run() error {
	fmt.Printf("Texted %s this message: %s%s\n", *textNumber, *message,
		*signature)
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
