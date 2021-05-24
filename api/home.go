package api

import "net/http"

// Home ...
func (rt *Router) Home(w http.ResponseWriter, r *http.Request) {

	resp := response{
		code: http.StatusOK,
		Data: "This is a great time",
	}

	resp.serveJSON(w)
}
