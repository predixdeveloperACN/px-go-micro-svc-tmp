package rest

import (
	"os"
	"time"
	"fmt"

	"encoding/json"
	"net/http/pprof"
	"errors"
	"log"

	swagger "github.com/predixdeveloperACN/swagger-ui"
	"github.com/gorilla/mux"
)

var startTime time.Time

func GetPredixSpace() (space string) {
	var v map[string]string

	vcapApp := os.Getenv("VCAP_APPLICATION")
	if debug {
		log.Println(fmt.Sprintf("DBG-> Vcap: %v", vcapApp))
	}

	err := json.Unmarshal([]byte(vcapApp), &v)
	if err != nil {
		log.Println(errors.New(fmt.Sprint("ERROR: Could not convert Vcap Services json data")))
	}

	if (v != nil) {
		space = v["space_name"]
		if debug {
			log.Println(fmt.Sprintf("DBG-> v: %v", v))
		}
	} else {
		space = "unknown"
	}

	return
}

func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

func getRouter() (router *mux.Router) {
	startTime = time.Now().UTC()

	log.Println("Configuring Base REST Server Interface at ", startTime.Format(time.RFC850))

	// Register a handler for each route pattern
	router = mux.NewRouter()

	// Custom REST handlers
	for _, route := range routes {
		router.Methods(route.Method).HandlerFunc(route.HandlerFunc).Path(route.Pattern).Name(route.Name)
	}

	// Add a trivial handler for INFO
	attachProfiler(router)

	// attach swagger documentation api
	swagger.AttachSwaggerUI(router, base_path)

	return
}

