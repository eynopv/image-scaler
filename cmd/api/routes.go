package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/image", app.uploadImageHandler)
	router.HandlerFunc(http.MethodGet, "/image/:id", app.downloadImageHandler)

	return router
}
