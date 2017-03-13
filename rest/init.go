package rest

import (
	"os"
	"strconv"
)

var oauthDisabled, debug bool

func init() {
	// get debug config
	var err error
	debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		// if local
		debug = true
	}

	oauthDisabled, err = strconv.ParseBool(os.Getenv("OAUTH2_DISABLED"))
	if err != nil {
		// if local
		oauthDisabled = true
	}
}