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
	level       int // the level at which to start args parsing
	Name        string
	Usage       string // the usage string
	Description string
	FlagSet     *flag.FlagSet
	CommandSet  map[string]*Command
	Execute     func(command *Command)
}

// NewCommand creates a basic new command.
func NewCommand(name string, handling flag.ErrorHandling) *Command {
	cmd := &Command{
		level:       0,
		Name:        name,
		Description: "",
		FlagSet:     flag.NewFlagSet(name, handling),
		CommandSet:  make(map[string]*Command),
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

// PrintHelp prints the command help to the console.
func (c Command) PrintHelp() {
	fmt.Printf("%s", c.Help())
	c.FlagSet.PrintDefaults()
}

// Add appends a command to the command set.
func (c *Command) Add(cmd *Command) error {
	// Fatal if the command already exists
	if _, ok := c.CommandSet[cmd.Name]; ok {
		return fmt.Errorf("cannot add command %s already exists", cmd.Name)
	}
	cmd.level = c.level + 1
	c.CommandSet[cmd.Name] = cmd
	return nil
}

// AddCommands appends multiple commands to the command set.
func (c *Command) AddCommands(cmds []*Command) error {
	for _, cmd := range cmds {
		err := c.Add(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}
