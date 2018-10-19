package cmdline

import (
	"flag"
	"fmt"
	"strings"
)

type Runner interface {
	Run(args []string) error
}

type Cmd struct {
	name   string
	subs   []*Cmd
	runner Runner
	flags  *flag.FlagSet
	args   []*Arg
}

func NewCmd(name string, runner Runner, args []*Arg, flags *flag.FlagSet) *Cmd {
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
	if len(c.args) < 1 {
		return ""
	}
	var argNames []string
	var argDescriptions []string
	for _, arg := range c.args {
		argNames = append(argNames, fmt.Sprintf("[%s]", arg.name))
		argDescriptions = append(argDescriptions, fmt.Sprintf("%s: %s",
			arg.name, arg.desc))
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

type Arg struct {
	name string
	desc string
}

func NewArg(name, desc string) *Arg {
	return &Arg{
		name: name,
		desc: desc,
	}
}
