package cmdline

import (
	"errors"
	"fmt"
	"os"
)

var (
	UsageErr  error = errors.New("Incorrect usage")
	runnerErr error = errors.New("Cannot run command without a runner")
	defineErr error = errors.New("Cannot define command with both args and subcommands")
)

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
	if len(c.subs) > 0 && c.args.Len() > 0 {
		return defineErr
	}
	if len(c.subs) > 0 {
		return UsageErr
	}
	if c.runner == nil {
		return runnerErr
	}
	if len(args) < c.args.Len() {
		return UsageErr
	}
	if c.flags != nil {
		if err := c.flags.Parse(args); err != nil {
			return err
		}
		args = c.flags.Args()
	}
	if c.args != nil {
		if err := c.args.Parse(args); err != nil {
			return err
		}
	}
	return c.runner.Run()
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
