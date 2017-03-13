package postgres

import (
	"os"
	"strconv"
)

var debug bool

func init() {

	// get debug config
	var err error
	debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		// if local
		debug = true
	}
}
