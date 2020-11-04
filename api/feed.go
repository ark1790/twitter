package api

import (
	"net/http"
)

func (rt *Router) GetFeeds(w http.ResponseWriter, r *http.Request) {
	usr := getAuthUser(r)

	twts, err := rt.feedRepo.List(usr)
	if err != nil {
		panic(newAPIError("Internal Server Error", errInternalServer, err))
	}
	resp := response{
		code: http.StatusOK,
		Data: twts,
	}

	resp.serveJSON(w)

}
