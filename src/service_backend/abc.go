package main

import (
	"fmt"
	"net/http"
)
// Abc aaa
func Abc() string{
	return "ABC function from abc.go"
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, Abc())
}