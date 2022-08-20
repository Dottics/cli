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
