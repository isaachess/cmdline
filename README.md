# CmdLine

CmdLine is a simple Go library to assist in creating command-line (CLI) tools in Go.
It allows you to declaritively specify your commands, flags, required arguments,
etc. It supports sub-commands, sub-command-specific flags and args, usage
messages, and more.

# Example

This contrived example has a cmdline tool called `phone` that will call and text
people for you.

It has two sub-commands: `phone call` and `phone text`. You can find it in
`examples/example.go`. To build, simply:

`go build -o phone`

```go
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
```

From this example, let's explore the automatic usage messages a little.

```
$ ./phone
Usage: phone [subcommand]

Subcommands:
call
text

```

```

$ ./phone call
Usage: call [number]

Args:
number: (123-456-7890) the phone number to call

Flags:
speaker: (bool) whether the speaker phone should be used

```

```

$ ./phone call 555-123-4567
Dialing 555-123-4567...

$ ./phone call -speaker 555-123-4567
Dialing 555-123-4567 with speaker phone...

```

```

$ ./phone text
Usage: text [number] [message]

Args:
number: (123-456-7890) the phone number to call
message: (string) the message to send

Flags:
signature: (string) signature to use at the end of the message
```

For a more complete (but still contrived) example, please see the `examples`
folder.

# Usage

## Cmd

The `Cmd` type represents a command, and is initialized with:

`cmdline.NewCmd(name string, runner Runner, args []*Arg, flags *flag.FlagSet)`

The `name` of the command is how you will invoke the command. For the main
parent command this isn't used, but for all subcommands the `name` element is
used to identify which command to execute.

## cmdline.Runner interface

The cmdline.Runner interface is:

```go
type Runner interface {
	Run(args []string) error
}
```

The `Run` function will be called when the command is invoked. It is passed the
args specific to this command, meaning it strips off the command name and any
parsed flags.

For example, suppose you have a command `phone` with two subcommands: `text` or
`call`. And suppose call requires you to supply, as arguments, the phone number
to call. Suppose it also allows a bool flag to specify whether speaker phone
should be on. This command would look like:

`phone call -speaker 555-326-1234`

In this case, the `call` command's `Run` function is called. The only argument
passed is `555-326-1234` as it is the only arg for this subcommand.

Any error returned by the `Run` function is printed at the top of a usage
statement to the user.

## cmdline.Arg

Args are initialized with:

`cmdline.NewArg(name, desc string) *Arg`

A list of args is registered with the command on creation. The arg order
matters: it is used to build a convenient usage message. All args are required
by the user, so `cmdline` will ensure the correct number are passed. If not, a
usage statement is printed.

## Flags

`cmdline` uses Go's built-in flag package. For each command, you may supply a
`*flag.FlagSet` specific for that command/subcommand. (You may provide `nil` if
no flags are needed.) The flags will be parsed prior to when the `Run` method
for the runner is called.

Flags are included in the usage statement printed for each command.
