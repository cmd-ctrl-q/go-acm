package main

import "net/http"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page"))
}
