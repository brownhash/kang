package core

import (
	"fmt"
	"time"
	"os/user"

	"github.com/brownhash/golog"
	"github.com/gofrs/uuid"
)

func GenerateSession() Session {
	uid, err := uuid.NewV4()
	if err != nil {
		golog.Error(err.Error())
	}

	user, err := user.Current()
	if err != nil {
		golog.Error(err.Error())
	}

	session := Session{
		Id: fmt.Sprintf("%v", uid),
		Started: time.Now().Local(),
		User: user.Username,
	}

	return session
}

type Session struct {
	Id 			string
	Started 	time.Time
	User		string 		
}