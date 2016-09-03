package main

import "net/http"

// Route models the components of a REST endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes acts to hold the routes of the web app, implemented below
type Routes []Route

var routes = Routes{
	Route{
		"ObjCreate",
		"POST",
		"/store",
		ObjCreate,
	},
	Route{
		"ObjRetrieve",
		"GET",
		"/retrieve",
		ObjRetrieve,
	},
}
