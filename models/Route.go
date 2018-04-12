package models

import "net/http"

//Struct definition of a route

type Route struct{
	Name 	string
	Method	string
	Pattern	string
	HandlerFunction http.HandlerFunc
}

type Routes []Route