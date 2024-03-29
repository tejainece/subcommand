package subcommands

import (
	"context"
	"flag"
)

// A Command represents a single command.
type Command interface {
	// Name returns the name of the command.
	Name() string

	// Synopsis returns a short string (less than one line) describing the command.
	Synopsis() string

	Description() string

	// Usage returns a long string explaining the command and giving usage
	// information.
	Usage() string

	Footer() string

	// SetFlags adds the flags for this command to the specified set.
	SetFlags(*flag.FlagSet)

	// Execute executes the command and returns an ExitStatus.
	Execute(commander *Commander, ctx context.Context, f *flag.FlagSet, args ...interface{}) ExitStatus
}

// An ExitStatus represents a Posix exit status that a subcommand
// expects to be returned to the shell.
type ExitStatus int

const (
	ExitSuccess ExitStatus = iota
	ExitFailure
	ExitUsageError
)

type CommandMixin struct {}

func (CommandMixin) Footer() string {return "" }

func (CommandMixin) SetFlags(f *flag.FlagSet) {}