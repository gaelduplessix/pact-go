package command

import (
	"os"
	"testing"
)

func init() {
	// Set CLI flags to simulate real
	os.Args = []string{"version"}
}

func Test_RootCommand(t *testing.T) {
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	Execute()
}