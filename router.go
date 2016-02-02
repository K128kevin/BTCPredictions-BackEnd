package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range UserManagement {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	for _, route := range TradeSimulator {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	for _, route := range DataAPIs {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var root = "/api"

// REST
var DataAPIs = Routes {
	// define apis for prices and predictions here
}

// REST
var UserManagement = Routes {
	// define each user management endpoint in here
}

// JSON-RPC
var TradeSimulator = Routes {
	// define each trade simulator endpoint in here
}