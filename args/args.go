package args

import (
	"errors"
	"strconv"
	"time"
)

var mismatchErr = errors.New(
	"Len of args to parse does not match len of registered args")

type argParser interface {
	Parse(arg string) error
}

type arg struct {
	parser argParser
	name   string
	desc   string
}

type ArgSet struct {
	args   []*arg
	parsed bool
}

func NewArgSet() *ArgSet {
	return &ArgSet{}
}

func (a *ArgSet) Parse(args []string) error {
	if len(args) != len(a.args) {
		return mismatchErr
	}
	for i, arg := range args {
		if err := a.args[i].parser.Parse(arg); err != nil {
			return err
		}
	}
	a.parsed = true
	return nil
}

func (a *ArgSet) addArg(name, desc string, par argParser) {
	a.args = append(a.args, &arg{par, name, desc})
}

func (a *ArgSet) Bool(name, desc string) *bool {
	par := NewBoolParser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Duration(name, desc string) *time.Duration {
	par := NewDurationParser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Float64(name, desc string) *float64 {
	par := NewFloat64Parser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Int(name, desc string) *int {
	par := NewIntParser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Int64(name, desc string) *int64 {
	par := NewInt64Parser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) String(name, desc string) *string {
	par := NewStringParser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Uint(name, desc string) *uint {
	par := NewUintParser()
	a.addArg(name, desc, par)
	return &par.v
}

func (a *ArgSet) Uint64(name, desc string) *uint64 {
	par := NewUint64Parser()
	a.addArg(name, desc, par)
	return &par.v
}

// ArgList returns the list of names of the args
func (a *ArgSet) Names() []string {
	var list []string
	for _, arg := range a.args {
		list = append(list, arg.name)
	}
	return list
}

func (a *ArgSet) Desc() []string {
	var list []string
	for _, arg := range a.args {
		list = append(list, arg.desc)
	}
	return list
}

func (a *ArgSet) Len() int {
	if a == nil {
		return 0
	}
	return len(a.args)
}

func (a *ArgSet) Parsed() bool {
	return a.parsed
}

type boolParser struct {
	v bool
}

func NewBoolParser() *boolParser {
	return &boolParser{}
}

func (par *boolParser) Parse(arg string) error {
	v, err := strconv.ParseBool(arg)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}

type durationParser struct {
	v time.Duration
}

func NewDurationParser() *durationParser {
	return &durationParser{}
}

func (par *durationParser) Parse(arg string) error {
	v, err := time.ParseDuration(arg)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}

type float64Parser struct {
	v float64
}

func NewFloat64Parser() *float64Parser {
	return &float64Parser{}
}

func (par *float64Parser) Parse(arg string) error {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}

type intParser struct {
	v int
}

func NewIntParser() *intParser {
	return &intParser{}
}

func (par *intParser) Parse(arg string) error {
	v, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}

type int64Parser struct {
	v int64
}

func NewInt64Parser() *int64Parser {
	return &int64Parser{}
}

func (par *int64Parser) Parse(arg string) error {
	v, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}

type stringParser struct {
	v string
}

func NewStringParser() *stringParser {
	return &stringParser{}
}

func (par *stringParser) Parse(arg string) error {
	par.v = arg
	return nil
}

type uintParser struct {
	v uint
}

func NewUintParser() *uintParser {
	return &uintParser{}
}

func (par *uintParser) Parse(arg string) error {
	v, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return err
	}
	par.v = uint(v)
	return nil
}

type uint64Parser struct {
	v uint64
}

func NewUint64Parser() *uint64Parser {
	return &uint64Parser{}
}

func (par *uint64Parser) Parse(arg string) error {
	v, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return err
	}
	par.v = v
	return nil
}
