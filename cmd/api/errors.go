package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path:%s  errors:%s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path:%s  errors:%s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("status not Found error: %s path:%s  errors:%s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusNotFound, err.Error())

}

func (app *application) ConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("status conflict error: %s path:%s  errors:%s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())

}
