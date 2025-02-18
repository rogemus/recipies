package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
		body   = http.StatusText(http.StatusInternalServerError)
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)

	if app.debugMode {
		body = fmt.Sprintf("%s\n%s", err, trace)
	}

	http.Error(w, body, http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.tmplCache[page]

	if !ok {
		err := fmt.Errorf("template: template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)

	if !ok {
		return false
	}

	return isAuthenticated
}

func (app *application) sessionUserId(r *http.Request) int {
	id := app.sessionManager.GetInt(r.Context(), userIdSessionKey)
	return id
}

func (app *application) sessionUserName(r *http.Request) string {
	name := app.sessionManager.GetString(r.Context(), userNameSessionKey)
	return name
}
