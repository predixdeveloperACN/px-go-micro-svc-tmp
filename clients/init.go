package clients

import (
	"os"
	"strconv"
)

var cupsName string
var debug bool

func init() {
	cupsName = os.Getenv("CUPS_NAME")

	// get debug config
	var err error
	debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		// if local
		debug = true
	}
}

func CupsName() string {
	return cupsName
}