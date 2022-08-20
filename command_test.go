package cli

import (
	"flag"
	"fmt"
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
	if (*cmd).Help() != helpString {
		t.Errorf("expected help string '%s' got '%s'", helpString, (*cmd).Help())
	}
}
