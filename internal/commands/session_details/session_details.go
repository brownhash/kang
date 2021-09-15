package session_details

import (
	"fmt"
	"github.com/brownhash/golog"
	s "github.com/brownhash/session_terraform/internal/session"
)

func Run(session s.Session) int {
	golog.Warn("Current session details:")
	golog.Warn(fmt.Sprintf("ID: %s", session.Id))
	golog.Warn(fmt.Sprintf("Started by: %s", session.User))
	golog.Warn(fmt.Sprintf("Started at: %s", session.Started))

	return 0
}

func Help() string {
	return "Provides a detailed log of current session, displaying its Unique Id, the user who started it and when."
}

func Synopsis() string {
	return "Current session details."
}