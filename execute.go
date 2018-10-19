package cmdline

import (
	"errors"
	"fmt"
	"os"
)

var UsageErr error = errors.New("Incorrect usage")

func Execute(c *Cmd) error {
	return executeWithPrint(c, os.Args[1:], true)
}

func executeWithPrint(c *Cmd, args []string, shouldPrint bool) error {
	c, args = findCmdAndArgs(c, args)
	err := execute(c, args)
	if err != nil && shouldPrint {
		fmt.Println(c.Usage(err))
	}
	return err
}

func execute(c *Cmd, args []string) error {
	if len(c.subs) > 0 && len(c.args) > 0 {
		return errors.New("Command definition error: Cannot define command with both args and subcommands")
	}
	if len(c.subs) > 0 {
		return UsageErr
	}
	if c.runner == nil {
		return UsageErr
	}
	if c.flags != nil {
		if err := c.flags.Parse(args); err != nil {
			return err
		}
		args = c.flags.Args()
	}
	if len(args) < len(c.args) {
		return UsageErr
	}
	return c.runner.Run(args)
}

func findCmdAndArgs(c *Cmd, args []string) (*Cmd, []string) {
	if subCmd := findSubCommand(c, args); subCmd != nil {
		return findCmdAndArgs(subCmd, args[1:])
	}
	return c, args
}

func findSubCommand(c *Cmd, args []string) *Cmd {
	if len(args) < 1 {
		return nil
	}
	for _, subCmd := range c.subs {
		if subCmd.name == args[0] {
			return subCmd
		}
	}
	return nil
}
