package main

import (
	"testing"
)


func ListenAndServe(addr string, handler Handler) error


type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func TestgetConferenceDetails(t *testing.T) {
t.Run()
}