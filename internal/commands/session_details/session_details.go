package session_details

import (
	"fmt"
	"github.com/brownhash/golog"
	"github.com/brownhash/session_terraform/internal/core"
)

func Run(session core.Session) int {
	golog.Warn("Current session details:")
	golog.Warn(fmt.Sprintf("ID: %s", session.Id))
	golog.Warn(fmt.Sprintf("Started by: %s", session.User))
	golog.Warn(fmt.Sprintf("Started at: %s", session.Started))

	return 0
}

func Help() string {
	return "\n\tProvides a detailed log of current session, displaying its Unique Id, the user who started it and when.\n"
}

func Synopsis() string {
	return "Current session details."
}