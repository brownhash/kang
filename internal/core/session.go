package core

import (
	"fmt"
	"os/user"
	"strings"
	"time"

	"github.com/brownhash/golog"
	"github.com/brownhash/kang/config"
	badger "github.com/dgraph-io/badger/v3"
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
		Id:      fmt.Sprintf("%v", uid),
		Started: time.Now().Local(),
		User:    user.Username,
	}

	return session
}

func ReadSessions() error {
	db, err := badger.Open(badger.DefaultOptions(config.SetupPath))
	if err != nil {
		return err
	}

	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
		  item := it.Item()
		  k := item.Key()

		  err := item.Value(func(v []byte) error {
			golog.Warn(fmt.Sprintf("%v: %v", k, v))
			return nil
		  })

		  if err != nil {
			return err
		  }
		}

		return nil
	  })

	return err
}

type Session struct {
	Id        string
	Started   time.Time
	User      string
	Command   string
	Arguments []string
}

func (s *Session) Save(exitStatus int) error {
	sessionMessage := fmt.Sprintf("Started by %s, at %v to run `%s` command with `%s` arguments. Exit status [%d].", s.User, s.Started, s.Command, strings.Join(s.Arguments, " "), exitStatus)
	golog.Debug(sessionMessage)

	db, err := badger.Open(badger.DefaultOptions(config.SetupPath))
	if err != nil {
		return err
	}

	txn := db.NewTransaction(true)

	err = txn.Set([]byte(s.Id), []byte(sessionMessage))
	if err != nil {
		return err
	}

	err = txn.Commit()

	if err != nil {
		return err
	}

	defer func() {
		txn.Discard()
		db.Close()
	}()

	return err
}
