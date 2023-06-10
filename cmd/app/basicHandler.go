package main

import (
	"errors"
	"github/valentedev/httpserver-go/internal/data"
	"net/http"
)

func (app *application) getBasic(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	thread, err := app.models.Threads.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"thread": thread}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
