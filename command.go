package cli

import (
	"errors"
	"flag"
	"fmt"
)

var help bool

// ErrorNotEqual is a function used for testing to test whether two errors
// are not equal and returns false if the errors are the same else true if
// the error are not equal.
func ErrorNotEqual(err1, err2 error) bool {
	if err1 == err2 {
		return false
	}
	if err1 != nil && err2 != nil {
		if err1.Error() != err2.Error() {
			return true
		} else {
			return false
		}
	}
	// this means one is nil and the other not
	return true
}

// WIP is a placeholder function used when create a new command.
func WIP(cmd *Command) error {
	fmt.Printf("\n\n** WIP **\n\n")
	return nil
}

// isCommand checks if the first element in the slice of strings from os.Args
// is a command, then returns the command, otherwise if the first element is a
// flag then it returns a zero string.
func isCommand(args []string) string {
	if len(args) == 0 {
		return ""
	}
	v := args[0]
	if v == "" { // although os.Args will never return a zero string
		return v
	}
	if string(v[0]) == "-" {
		// if v does start with '-' then it is a flag
		return ""
	} else {
		// if v does not start with '-' then it is a command
		return v
	}
}

// Command is a struct.
type Command struct {
	level       int // the level at which to start args parsing
	Name        string
	Usage       string // the usage string
	Description string
	FlagSet     *flag.FlagSet
	CommandSet  Commands
	Execute     func(command *Command) error
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
	cmd.FlagSet.BoolVar(&help, "help", false, "To get help information for this command.")
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
func (c *Command) PrintHelp() {
	fmt.Printf("%s", c.Help())
	c.FlagSet.PrintDefaults()
	fmt.Printf("%s", c.CommandSet.Help())
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
func (c *Command) AddCommands(xc []*Command) error {
	for _, cmd := range xc {
		err := c.Add(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run is used to run the command, this could be to call a sub-command or
// to execute the command pointed to.
func (c *Command) Run(args []string) error {
	// if there are no args print help
	if len(args) == 0 {
		c.PrintHelp()
		return errors.New("invalid operation: required args length > 0")
	}

	// for this command parse the args
	err := c.Init(args)
	if err != nil {
		return err
	}

	command := isCommand(args)
	sub, ok := c.CommandSet[command]
	if ok {
		return sub.Run(args[1:])
	} else {
		return c.Execute(c)
	}
}
