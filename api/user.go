package api

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/ark1790/alpha/errors"
	"github.com/ark1790/alpha/model"
)

type createUserPld struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Private  bool   `json:"private"`
}

func (c *createUserPld) validate() *validationError {
	c.Username = strings.TrimSpace(c.Username)

	errV := validationError{}
	if c.Name == "" {
		errV.add("name", "is required")
	}
	if c.Username == "" {
		errV.add("username", "is required")
	}

	if len(errV) > 0 {
		return &errV
	}

	return nil
}

// CreateUser ...
func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	body := createUserPld{}
	if err := parseBody(r, &body); err != nil {
		panic(newAPIError("Unable to parse body", errBadRequest, err))
	}

	if err := body.validate(); err != nil {
		panic(newAPIError("Invalid data", errInvalidData, err))
	}

	user := &model.User{
		Name:     body.Name,
		Username: body.Username,
		Private:  body.Private,
	}

	if err := rt.userRepo.Create(user); err != nil {
		if err == errors.ErrDuplicateKey {
			vErr := validationError{}
			vErr.add("username", "is not unique")
			panic(newAPIError("Invalid data", errEntityNotUnique, &vErr))
		}

		panic(newAPIError("Internal Server Error", errInternalServer, err))
	}

	resp := response{
		code: http.StatusOK,
		Data: user,
	}

	resp.serveJSON(w)
}

type loginPld struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *loginPld) validate() *validationError {
	errV := validationError{}
	if l.Username == "" {
		errV.add("username", "is required")
	}
	if l.Password == "" {
		errV.add("password", "is required")
	}

	if len(errV) > 0 {
		return &errV
	}

	return nil
}

func (rt *Router) Login(w http.ResponseWriter, r *http.Request) {
	body := loginPld{}
	if err := parseBody(r, &body); err != nil {
		panic(newAPIError("Unable to parse body", errBadRequest, err))
	}

	if err := body.validate(); err != nil {
		panic(newAPIError("Invalid data", errInvalidData, err))
	}

	usr, err := rt.userRepo.Fetch(body.Username)
	if err != nil {
		panic(newAPIError("DB failed", errInternalServer, err))
	}
	if usr == nil {
		panic(newAPIError("User not found", errUserNotFound, nil))
	}

	data := []byte(body.Username + ":" + body.Password)
	token := base64.StdEncoding.EncodeToString(data)

	vData := []byte(usr.Username + ":" + usr.Username)
	vToken := base64.StdEncoding.EncodeToString(vData)

	if token != vToken {
		panic(newAPIError("Unauthorized", errUnAuthorized, nil))
	}

	resp := response{
		code: http.StatusOK,
		Data: object{
			"token": vToken,
		},
	}

	resp.serveJSON(w)

}

func (rt *Router) GetMe(w http.ResponseWriter, r *http.Request) {
	uName := getAuthUser(r)

	usr, err := rt.userRepo.Fetch(uName)
	if err != nil {
		panic(newAPIError("DB failed", errInternalServer, err))
	}
	if usr == nil {
		panic(newAPIError("User not found", errUserNotFound, nil))
	}

	cFlg, cFlw, err := rt.followRepo.Count(uName)
	if err != nil {
		panic(newAPIError("DB failed", errInternalServer, err))
	}
	resp := response{
		code: http.StatusOK,
		Data: object{
			"user":      usr,
			"following": cFlg,
			"follower":  cFlw,
		},
	}

	resp.serveJSON(w)

}
