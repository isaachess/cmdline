package cmdline

import (
	"cmdline/args"
	"flag"
	"testing"
)

const (
	cmd1Name   = "1"
	cmd11Name  = "11"
	cmd111Name = "111"
	cmd112Name = "112"
	cmd12Name  = "12"
	cmd13Name  = "13"
	cmd14Name  = "14"
)

type testRunner struct {
	run bool
	err error
}

func (c *testRunner) Run() error {
	c.run = true
	return c.err
}

func TestSubcommandIsRun(t *testing.T) {
	test := newTest(t)
	test.Execute([]string{cmd11Name, cmd111Name, "bobby", "jones"})
	test.AssertRun(cmd111Name)
}

func TestParentCommandWontRunWithoutSubSpecified(t *testing.T) {
	test := newTest(t)
	test.AssertUsageError([]string{cmd11Name})
}

func TestUsageErrIfNoRunner(t *testing.T) {
	test := newTest(t)
	test.Execute([]string{cmd13Name})
	test.AssertNoRuns()
	test.AssertExecuteError(runnerErr)
}

func TestUsageErrIfTooFewArgs(t *testing.T) {
	test := newTest(t)
	test.AssertUsageError([]string{cmd11Name, cmd111Name})

	test = newTest(t)
	test.AssertUsageError([]string{cmd11Name, cmd111Name, "bobby"})
}

func TestErrorIfSubsAndArgsDefined(t *testing.T) {
	test := newTest(t)
	test.Execute([]string{cmd14Name})
	test.AssertExecuteError(defineErr)
	test.AssertNoRuns()
}

type test struct {
	cmd  *Cmd
	t    *testing.T
	cmds map[string]*Cmd
	err  error
}

func newTest(t *testing.T) *test {
	var (
		cmd111Flags = flag.NewFlagSet(cmd111Name, flag.ExitOnError)
		cmd112Flags = flag.NewFlagSet(cmd112Name, flag.ExitOnError)
		cmd11Flags  = flag.NewFlagSet(cmd11Name, flag.ExitOnError)

		cmd12Flags = flag.NewFlagSet(cmd12Name, flag.ExitOnError)

		cmd1Flags = flag.NewFlagSet(cmd1Name, flag.ExitOnError)
	)

	var (
		cmd111Args = args.NewArgSet()
		_          = cmd111Args.String("bobby", "a name bobby")
		_          = cmd111Args.String("jones", "a name jones")

		cmd14Args = args.NewArgSet()
		_         = cmd14Args.String("badarg", "this shouldn't happen")
	)

	cmd111 := NewCmd(cmd111Name, &testRunner{}, cmd111Args, cmd111Flags)
	cmd112 := NewCmd(cmd112Name, &testRunner{}, nil, cmd112Flags)
	cmd11 := NewCmd(cmd11Name, &testRunner{}, nil, cmd11Flags)
	cmd11.RegisterSub(cmd111)
	cmd11.RegisterSub(cmd112)

	cmd12 := NewCmd(cmd12Name, &testRunner{}, nil, cmd12Flags)
	cmd13 := NewCmd(cmd13Name, nil, nil, nil)
	cmd14 := NewCmd(cmd14Name, &testRunner{}, cmd14Args, nil)
	cmd14.RegisterSub(NewCmd("BADSUB", &testRunner{}, nil, nil))

	cmd1 := NewCmd(cmd1Name, &testRunner{}, nil, cmd1Flags)
	cmd1.RegisterSub(cmd11)
	cmd1.RegisterSub(cmd12)
	cmd1.RegisterSub(cmd13)
	cmd1.RegisterSub(cmd14)
	return &test{
		cmd: cmd1,
		t:   t,
		cmds: map[string]*Cmd{
			cmd1Name:   cmd1,
			cmd11Name:  cmd11,
			cmd111Name: cmd111,
			cmd112Name: cmd112,
			cmd12Name:  cmd12,
			cmd13Name:  cmd13,
			cmd14Name:  cmd14,
		},
	}
}

func (t *test) Execute(args []string) {
	if err := executeWithPrint(t.cmd, args, false); err != nil {
		t.err = err
	}
}

func (t *test) AssertUsageError(args []string) {
	t.Execute(args)
	t.AssertNoRuns()
	t.AssertExecuteError(UsageErr)
}

func (t *test) AssertExecuteError(err error) {
	if err != t.err {
		t.t.Error("Errors do not match")
	}
}

func (t *test) AssertRun(cmdName string) {
	for name, cmd := range t.cmds {
		if name == cmdName {
			t.assertRun(name, cmd)
		} else {
			t.assertNotRun(name, cmd)
		}
	}
}

func (t *test) AssertNoRuns() {
	for name, cmd := range t.cmds {
		t.assertNotRun(name, cmd)
	}
}

func (t *test) assertRun(cmdName string, cmd *Cmd) {
	tRunner := cmd.runner.(*testRunner)
	if !tRunner.run {
		t.t.Errorf("%s was not run but should have", cmdName)
	}
	if cmd.flags != nil {
		if !cmd.flags.Parsed() {
			t.t.Errorf("%s: flags not parsed", cmdName)
		}
	}
	if cmd.args != nil {
		if !cmd.args.Parsed() {
			t.t.Errorf("%s: args not parsed", cmdName)
		}
	}
}

func (t *test) assertNotRun(cmdName string, cmd *Cmd) {
	if cmd.runner == nil {
		return
	}
	if cmd.runner.(*testRunner).run {
		t.t.Errorf("%s was run but should not have", cmdName)
	}
	if cmd.flags != nil {
		if cmd.flags.Parsed() {
			t.t.Errorf("%s: flags parsed", cmdName)
		}
	}
}
