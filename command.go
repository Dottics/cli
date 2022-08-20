package cli

import (
	"flag"
	"fmt"
)

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
