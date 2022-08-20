package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestNewCommand(t *testing.T) {
	cmd := NewCommand("get", flag.ExitOnError)

	if cmd.Name != "get" {
		t.Errorf("expected command name '%s' got '%s'", "get", cmd.Name)
	}
	if cmd.FlagSet == nil {
		t.Errorf("expected new command to have FlagSet got %v", nil)
	}
	if cmd.CommandSet == nil {
		t.Errorf("expected command set %v got %v", nil, cmd.CommandSet)
	}
	if len(cmd.CommandSet) != 0 {
		t.Errorf("expected command set to have length %d got %d", 0, len(cmd.CommandSet))
	}
}

func TestCommand_Help(t *testing.T) {
	cmd := NewCommand("get", flag.ExitOnError)
	cmd.Usage = "cli"
	cmd.Description = "get some info."
	helpString := fmt.Sprintf("Usage: %s %s\n\n%s\n\n", cmd.Usage, cmd.Name, cmd.Description)
	if cmd.Help() != helpString {
		t.Errorf("expected help string '%s' got '%s'", helpString, cmd.Help())
	}
}

func TestCommand_Init(t *testing.T) {
	// best practise to restore the global state to as before
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	var username string

	cmd := NewCommand("get", flag.ContinueOnError)
	cmd.FlagSet.StringVar(&username, "username", "", "username flag.")

	tt := []struct {
		name     string
		args     []string
		err      error
		username string
	}{
		{
			name:     "spaced flag: -username james",
			args:     []string{"cli", "get", "-username", "james"},
			err:      nil,
			username: "james",
		},
		{
			name:     "equal flag: -username=james",
			args:     []string{"cli", "get", "-username=james"},
			err:      nil,
			username: "james",
		},
		{
			name:     "blank flag: -username",
			args:     []string{"cli", "get", "-username"},
			err:      errors.New("flag needs an argument: -username"),
			username: "",
		},
	}

	for _, tc := range tt {
		username = ""
		t.Run(tc.name, func(t *testing.T) {
			err := cmd.Init(tc.args[2:])
			if ErrorNotEqual(tc.err, err) {
				t.Errorf("expected err '%v' got '%v'", tc.err, err)
			}
			if username != tc.username {
				t.Errorf("expected flag username '%s' got '%s'", tc.username, username)
			}
		})
	}
}

func TestCommand_Add(t *testing.T) {
	c1 := NewCommand("add", flag.ExitOnError)
	c2 := NewCommand("user", flag.ExitOnError)
	c3 := NewCommand("location", flag.ExitOnError)
	c4 := NewCommand("user", flag.ExitOnError)

	err := c1.Add(c2)
	if err != nil {
		t.Errorf("expected error on first add %v got %v", nil, err)
	}
	err = c1.Add(c3)
	if err != nil {
		t.Errorf("expected error on first add %v got %v", nil, err)
	}
	err = c1.Add(c4)
	if err == nil {
		t.Errorf("expected error on second add to not be nil")
	}
}

func TestCommand_Add_level_cascade(t *testing.T) {
	c1 := NewCommand("add", flag.ExitOnError)
	c2 := NewCommand("user", flag.ExitOnError)
	c3 := NewCommand("location", flag.ExitOnError)
	c4 := NewCommand("address", flag.ExitOnError)

	if c1.level != 0 {
		t.Errorf("expected c1 to have level %d got %d", 0, c1.level)
	}
	if c2.level != 0 {
		t.Errorf("expected c1 to have level %d got %d", 0, c2.level)
	}
	if c3.level != 0 {
		t.Errorf("expected c1 to have level %d got %d", 0, c3.level)
	}
	if c4.level != 0 {
		t.Errorf("expected c1 to have level %d got %d", 0, c4.level)
	}

	err := c1.Add(c2)
	if err != nil {
		t.Errorf("expected error on first add %v got %v", nil, err)
	}
	err = c1.Add(c3)
	if err != nil {
		t.Errorf("expected error on first add %v got %v", nil, err)
	}

	if c1.level != 0 {
		t.Errorf("expected c1 to have level %d got %d", 0, c1.level)
	}
	if c2.level != 1 {
		t.Errorf("expected c1 to have level %d got %d", 1, c2.level)
	}
	if c3.level != 1 {
		t.Errorf("expected c1 to have level %d got %d", 1, c3.level)
	}

	err = c3.Add(c4)
	if err != nil {
		t.Errorf("expected error on first add %v got %v", nil, err)
	}
	if c4.level != 2 {
		t.Errorf("expected c1 to have level %d got %d", 2, c4.level)
	}
}

func TestCommand_AddCommands(t *testing.T) {
	tt := []struct {
		name string
		cmd  *Command
		cmds []*Command
		err  error
	}{
		{
			name: "no duplicates commands",
			cmd:  NewCommand("get", flag.ExitOnError),
			cmds: []*Command{
				NewCommand("user", flag.ExitOnError),
				NewCommand("task", flag.ExitOnError),
				NewCommand("contact", flag.ExitOnError),
			},
			err: nil,
		},
		{
			name: "duplicate commands",
			cmd:  NewCommand("get", flag.ExitOnError),
			cmds: []*Command{
				NewCommand("user", flag.ExitOnError),
				NewCommand("task", flag.ExitOnError),
				NewCommand("user", flag.ExitOnError),
			},
			err: fmt.Errorf("cannot add command user already exists"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.AddCommands(tc.cmds)
			if ErrorNotEqual(err, tc.err) {
				t.Errorf("expected error %v got %v", tc.err, err)
			}
		})
	}
}

func TestIsCommand(t *testing.T) {
	tt := []struct {
		name   string
		args   []string
		output string
	}{
		{
			name:   "no args",
			args:   []string{},
			output: "",
		},
		{
			name:   "command",
			args:   []string{"get", "-help"},
			output: "get",
		},
		{
			name:   "command",
			args:   []string{"get", "--help"},
			output: "get",
		},
		{
			name:   "command",
			args:   []string{"get", "--help", "subcommand"},
			output: "get",
		},
		{
			name:   "flags",
			args:   []string{"-help"},
			output: "",
		},
		{
			name:   "flags",
			args:   []string{"-help", "subcommand"},
			output: "",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := isCommand(tc.args)
			if o != tc.output {
				t.Errorf("expected command '%s' got '%s'", tc.output, o)
			}
		})
	}
}

func TestCommand_Run(t *testing.T) {
	tt := []struct {
		name string
	}{
		{
			name: "level 0 no args",
		},
		{
			name: "sub level no args",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
