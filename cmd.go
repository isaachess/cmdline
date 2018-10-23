package cmdline

import (
	"cmdline/args"
	"flag"
	"fmt"
	"strings"
)

type Runner interface {
	Run() error
}

type Cmd struct {
	name   string
	subs   []*Cmd
	runner Runner
	flags  *flag.FlagSet
	args   *args.ArgSet
}

func NewCmd(name string, runner Runner, args *args.ArgSet, flags *flag.FlagSet) *Cmd {
	return &Cmd{
		name:   name,
		runner: runner,
		flags:  flags,
		args:   args,
	}
}

func (c *Cmd) RegisterSub(cmd *Cmd) {
	c.subs = append(c.subs, cmd)
}

func (c *Cmd) RegisterSubs(cmds []*Cmd) {
	c.subs = append(c.subs, cmds...)
}

func (c *Cmd) Usage(err error) string {
	var errMsg string
	if err != nil && err != UsageErr {
		errMsg = fmt.Sprintf("Error: %s\n\n", err)
	}
	return fmt.Sprintf("%sUsage: %s%s%s%s", errMsg, c.name, c.argUsage(), c.subUsage(),
		c.flagUsage())
}

func (c *Cmd) subUsage() string {
	if len(c.subs) < 1 {
		return ""
	}
	var subNames []string
	for _, sub := range c.subs {
		subNames = append(subNames, sub.name)
	}
	return fmt.Sprintf(" [subcommand]\n\nSubcommands:\n%s",
		strings.Join(subNames, "\n"))
}

func (c *Cmd) argUsage() string {
	if c.args.Len() < 1 {
		return ""
	}
	var argNames = c.args.Names()
	var argDescriptions = c.args.Desc()
	for i, argName := range argNames {
		argNames[i] = fmt.Sprintf("[%s]", argName)
		argDescriptions[i] = fmt.Sprintf("%s: %s", argName, argDescriptions[i])
	}
	return fmt.Sprintf(" %s\n\nArgs:\n%s",
		strings.Join(argNames, " "), strings.Join(argDescriptions, "\n"))
}

func (c *Cmd) flagUsage() string {
	if c.flags == nil {
		return ""
	}
	var flags []string
	c.flags.VisitAll(func(f *flag.Flag) {
		flags = append(flags, fmt.Sprintf("%s: %s", f.Name, f.Usage))
	})
	return fmt.Sprintf("\n\nFlags:\n%s", strings.Join(flags, "\n"))
}
