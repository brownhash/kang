package session_details

import (
	"fmt"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/internal/core"
)

func Run(args []string) int {
	if len(args) > 0 {
		golog.Error(fmt.Sprintf("Unrecognised arguments %v. session accepts no arguments", args))
		return 1
	}

	golog.Info("Fetching existing session details")
	err := core.ReadSessions()

	if err != nil {
		return 1
	}

	return 0
}

func Help() string {
	return "\n\tProvides a detailed log of current session, displaying its Unique Id, the user who started it and when.\n"
}

func Synopsis() string {
	return "Current session details."
}