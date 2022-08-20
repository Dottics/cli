package cli

import (
	"flag"
	"fmt"
)

func ErrorNotEqual(err1, err2 error) bool {
	if err1 == err2 {
		return false
	}
	if err1 != nil && err2 != nil {
		if err1.Error() != err2.Error() {
			return true
		}
	}
	return false
}

// WIP is a placeholder function used when create a new command.
func WIP(cmd *Command) {
	fmt.Printf("WIP\n\n")
}

// Command is a struct
type Command struct {
	Name        string
	Usage       string
	Description string
	FlagSet     *flag.FlagSet
	//CommandSet  *CommandSet
	Execute func(command *Command)
}

// NewCommand creates a basic new command.
func NewCommand(name string, handling flag.ErrorHandling) *Command {
	cmd := &Command{
		Name:        name,
		Description: "",
		FlagSet:     flag.NewFlagSet(name, handling),
		Execute:     WIP,
	}
	return cmd
}

// Help is the method that prints the help description of the command to
// the Standard Output.
func (c *Command) Help() string {
	return fmt.Sprintf("Usage: %s %s\n\n%s\n\n", c.Usage, c.Name, c.Description)
}

// Init parses the command line args to the command's flags.
func (c *Command) Init(args []string) error {
	return c.FlagSet.Parse(args)
}
