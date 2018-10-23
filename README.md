# CmdLine

CmdLine is a simple Go library to assist in creating command-line (CLI) tools in
Go.  It allows you to declaritively specify your commands, flags, required
arguments, etc. It supports sub-commands, sub-command-specific flags and args,
usage messages, and more.

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

`cmdline.NewCmd(name string, runner Runner, args *args.ArgSet, flags *flag.FlagSet)`

The `name` of the command is how you will invoke the command. For the main
parent command this isn't used, but for all subcommands the `name` element is
used to identify which command to execute.

## Args

`cmdline` provides a parsable arg structure modeled on Go's built-in flags. You
can create an `ArgSet`, register args, and parse args. When args are parsed it
will populate the value pointers given when you register the args.

Args are initialized with:

`args.NewArgSet()`

You can then create args similar to flags:

```go
cmdCallArgs = args.NewArgSet()
callNumber  = cmdCallArgs.String("number",
    "(123-456-7890) the phone number to call")
```

In this example, `callNumber` is a `*string`, which will be populated when args
are parsed. Arg parsing happens automatically for you when you execute your cmd.

The _order_ you define args is significant. It will represent the expected order
of args when the cmd is run, and when printing usage.

If args cannot be parsed, or if the incorrect number of args is provided, a
usage message will be printed.

If you have no args for a command, you may pass `nil` to the cmd constructor.

## cmdline.Runner interface

The cmdline.Runner interface is:

```go
type Runner interface {
	Run() error
}
```

The `Run` function will be called when the command is invoked.

Any error returned by the `Run` function is printed at the top of a usage
statement to the user.

## Flags

`cmdline` uses Go's built-in flag package. For each command, you may supply a
`*flag.FlagSet` specific for that command/subcommand. (You may provide `nil` if
no flags are needed.) The flags will be parsed prior to when the `Run` method
for the runner is called.

Flags are included in the usage statement printed for each command.
