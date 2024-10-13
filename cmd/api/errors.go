package main

import (
	"net/http"
)

func (app *application) badRequestErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(w, r, http.StatusInternalServerError, "server encountered a problem")
}

func (app *application) notFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, "requested resource not found")
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	data := map[string]any{
		"error": message,
	}
	err := app.writeJson(w, status, data)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
}
