package main

import (
	"errors"
	"github/valentedev/httpserver-go/internal/data"
	"net/http"
)

func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.models.User.Get()

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"users": users}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
