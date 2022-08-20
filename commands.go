package cli

import "fmt"

// Commands is a slice of all commands for the command line tool
type Commands map[string]*Command

// Help formats the string that is printed to the os.StdOut when
// the --help flag is passed.
func (c Commands) Help() string {
	s := ""
	for _, cmd := range c {
		s += fmt.Sprintf("  %-10s  %s\n", cmd.Name, cmd.Description)
	}
	return s
}
