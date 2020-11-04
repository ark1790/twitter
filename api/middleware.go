package api

import (
	"context"
	"net/http"
	"strings"
)

const (
	ctxKeyUser string = "user"
)

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
