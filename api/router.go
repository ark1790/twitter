package api

import (
	"net/http"

	"github.com/ark1790/alpha/repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Router ...
type Router struct {
	*chi.Mux
	userRepo repo.User
}

// NewRouter ...
func NewRouter(ur repo.User) *Router {
	router := &Router{
		Mux:      chi.NewRouter(),
		userRepo: ur,
	}
	register(router)
	return router
}

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.DebugLevel)
}

func register(router *Router) {

	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	router.Use(recoverer)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Not Found", errURINotFound, nil)
		panic(err)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Method Not Allowed", errInvalidMethod, nil)
		resp := response{
			code:   http.StatusMethodNotAllowed,
			Errors: []apiError{*err},
		}
		resp.serveJSON(w)
	})

	router.Route("/", func(r chi.Router) {
		r.Mount("/users", userHandlers(router))
	})
}

func userHandlers(rt *Router) http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", rt.CreateUser)
		r.Post("/login", rt.Login)
	})

	return h
}
