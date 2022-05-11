package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	if status != 200 {
		app.logger.Println(message)
		w.WriteHeader(status)
	}
}

func (app *application) rateLimitExceedResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}
