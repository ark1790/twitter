package api

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func b64decode(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func getAuthToken(r *http.Request) string {
	aTkn := strings.TrimSpace(r.Header.Get("Authorization"))
	if strings.HasPrefix(aTkn, "Bearer ") {
		return strings.TrimPrefix(aTkn, "Bearer ")
	}
	return ""
}

// getAuthUser get authenticate user
func getAuthUser(r *http.Request) string {
	if usr := r.Context().Value(ctxKeyUser); usr != nil {
		return usr.(string)
	}
	return ""
}
