package postgres

import (
	"errors"
)

func PingPostgres() (err error) {

	if db_initialized {
		err = Database.Ping()
	} else {
		err = errors.New("Database not initialized")
	}

	return
}
