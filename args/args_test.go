package args

import (
	"testing"
	"time"
)

func TestArgSet(t *testing.T) {
	dur := "10s"
	args := NewArgSet()
	b := args.Bool("skipped", "this is how we skip it")
	d := args.Duration("how long", "this is how long we skip")
	f := args.Float64("pi", "what is pi")
	i := args.Int("and int", "make this 10")
	i64 := args.Int64("an int64", "make this 1000000")
	s := args.String("bobbyjones", "no idea")
	u := args.Uint("a uint", "some uint")
	u64 := args.Uint64("a uint64", "some uint64")
	if err := args.Parse([]string{
		"true",
		dur,
		"10.5",
		"10",
		"1000000",
		"bobby jones lived here",
		"15",
		"16",
	}); err != nil {
		t.Error("Parse failed", err)
	}
	expectedD, _ := time.ParseDuration(dur)
	if !*b {
		t.Error("Should be true")
	}
	if *d != expectedD {
		t.Error("Durations did not match")
	}
	if *f != 10.5 {
		t.Error("Floats did not match")
	}
	if *i != 10 {
		t.Error("Ints did not match")
	}
	if *i64 != 1000000 {
		t.Error("Int64s did not match")
	}
	if *s != "bobby jones lived here" {
		t.Error("Strings did not match")
	}
	if *u != 15 {
		t.Error("Uints did not match")
	}
	if *u64 != 16 {
		t.Error("Uint64s did not match")
	}
	if !args.Parsed() {
		t.Error("Should be marked parsed")
	}
}

func TestMismatchedArgLength(t *testing.T) {
	args := NewArgSet()
	args.Bool("hello", "world")
	args.Uint64("bye", "now")

	err := args.Parse([]string{"short"})
	if err != mismatchErr {
		t.Error("Should be mismatchErr")
	}
	err = args.Parse([]string{"too", "long", "now"})
	if err != mismatchErr {
		t.Error("Should be mismatchErr")
	}
}

func TestBoolParseErr(t *testing.T) {
	args := NewArgSet()
	args.Bool("hello", "world")
	if err := args.Parse([]string{"wat"}); err == nil {
		t.Error("Should err")
	}
}

func TestDurationParseErr(t *testing.T) {
	args := NewArgSet()
	args.Duration("hello", "world")
	if err := args.Parse([]string{"uhhh"}); err == nil {
		t.Error("Should err")
	}
}

func TestFloat64ParseErr(t *testing.T) {
	args := NewArgSet()
	args.Float64("hello", "world")
	if err := args.Parse([]string{"uhhh"}); err == nil {
		t.Error("Should err")
	}
}

func TestIntParseErr(t *testing.T) {
	args := NewArgSet()
	args.Int("hello", "world")
	if err := args.Parse([]string{"12.5"}); err == nil {
		t.Error("Should err")
	}
}

func TestInt64ParseErr(t *testing.T) {
	args := NewArgSet()
	args.Int64("hello", "world")
	if err := args.Parse([]string{"12.5"}); err == nil {
		t.Error("Should err")
	}
}

func TestUintParseErr(t *testing.T) {
	args := NewArgSet()
	args.Uint("hello", "world")
	if err := args.Parse([]string{"12.5"}); err == nil {
		t.Error("Should err")
	}
}

func TestUint64ParseErr(t *testing.T) {
	args := NewArgSet()
	args.Uint64("hello", "world")
	if err := args.Parse([]string{"12.5"}); err == nil {
		t.Error("Should err")
	}
}
