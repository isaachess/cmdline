package cmdline

import (
	"flag"
	"testing"
)

func TestCorrectCmdIsRun(t *testing.T) {
	var runBCalled bool
	bRun := func(args []string) error {
		runBCalled = true
		return nil
	}
	bFlags := flag.NewFlagSet("c0b", flag.PanicOnError)
	parent := NewCmd("c0", nil, nil)
	subA := NewCmd("c0a", nil, nil)
	subAA := NewCmd("c0aa")
	subB := NewCmd("c0b", bRun, bFlags)
	parent.RegisterSub(subA)
	parent.RegisterSub(subB)

	err := executeWithPrint(parent, []string{"c0b"}, false)
	if err != nil {
		t.Error(err)
	}
	if !runBCalled {
		t.Error("B should have been called")
	}
}

type test struct {
	c *Cmd
}

func buildTestCmd() *cmd {
	pass
	-copy
	-create
	-delete
}
