package api

import (
	"context"
	"net/http"
	"strings"
)

const (
	ctxKeyUser string = "user"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
	})
	// process the request...
}

func gatekeeper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := getAuthToken(r)
		if strings.TrimSpace(tkn) == "" {
			panic(newAPIError("Unauthorized", errUnAuthorized, nil))
		}

		dTokn, err := b64decode(tkn)
		if err != nil {
			panic(newAPIError("Unauthorized", errUnAuthorized, nil))
		}

		sTokn := strings.Split(dTokn, ":")
		if len(sTokn) < 2 {
			panic(newAPIError("Unauthorized", errUnAuthorized, nil))
		}

		uID := sTokn[0]
		if uID == "" {
			panic(newAPIError("Unauthorized", errUnAuthorized, nil))
		}

		r = r.WithContext(context.WithValue(r.Context(), ctxKeyUser, uID))
		next.ServeHTTP(w, r)
	})
}
