package cli

import (
	"fmt"
	"testing"
)

func TestCommands_Help(t *testing.T) {
	mc := Commands{
		"get": &Command{
			Name:        "get",
			Description: "Get all <entries>",
		},
		"adds": &Command{
			Name:        "adds",
			Description: "Add a new <entry>",
		},
	}

	expected := fmt.Sprintf("  get         Get all <entries>\n  adds        Add a new <entry>\n")
	helpString := mc.Help()
	if helpString != expected {
		t.Errorf("expected help string '%v' got '%v'", expected, helpString)
	}
}
