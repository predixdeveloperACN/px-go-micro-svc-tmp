package rest

import (
	"os"
	"net/http"
	"log"
	"github.com/gorilla/handlers"
	vcap "github.com/predixdeveloperACN/vcap-support"
)

const default_port = "8080"

var response_path string

func StartServer() {

	// Get Cloud Foundry assigned port
	port := os.Getenv("PORT")
	if port == "" {
		port = default_port
		log.Printf("Port number is not set using default: %s\n", port )
	}

	vcapAppMap, _ := vcap.LoadApplication()

	var uri string
	if len(vcapAppMap.URIs) > 0 {
		uri = vcapAppMap.URIs[0]
	}
	if uri == "" {
		response_path = "localhost:" + port + base_path
	} else {
		response_path = "https://" + uri + base_path
	}

	log.Println("Starting REST Server Interface for", uri)
	log.Println("response_path = %s", response_path)
	log.Println("Port = 0x%s", port)

	// get router object
	router := getRouter()

	// Start listening on the configured port.
	// ListenAndServe is not expected to return, so we wrap it in a log.Fatal
	// also include CORS handling
	err:= http.ListenAndServe(":"+port, handlers.CORS()(router))
	if(err != nil){
		log.Println(err)
	}
}
