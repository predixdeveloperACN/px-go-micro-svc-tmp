package rest

import (
	"net/http"
)

const base_path = "/api/v1/"

var routes = ConfiguredRoutes{
	Route{
		Name:        "Info",
		Method:      "GET",
		Pattern:     base_path + "info",
		HandlerFunc: HandleInfo},
	Route{
		Name:        "Ping",
		Method:      "GET",
		Pattern:     base_path + "ping",
		HandlerFunc: HandlePing},
}

type ConfiguredRoutes []Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}