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
	userRepo   repo.User
	followRepo repo.Follow
	tweetRepo  repo.Tweet
	feedRepo   repo.Feed
}

// NewRouter ...
func NewRouter(ur repo.User, fr repo.Follow, tr repo.Tweet, fdr repo.Feed) *Router {
	router := &Router{
		Mux:        chi.NewRouter(),
		userRepo:   ur,
		followRepo: fr,
		tweetRepo:  tr,
		feedRepo:   fdr,
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
	router.Use(enableCORS)

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
		r.Mount("/follows", followHandlers(router))
		r.Mount("/tweets", tweetHandlers(router))
		r.Mount("/feeds", feedHandlers(router))
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

func followHandlers(rt *Router) http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Use(gatekeeper)
		r.Post("/", rt.ToggleFollow)
	})

	return h
}
func tweetHandlers(rt *Router) http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Use(gatekeeper)
		r.Post("/", rt.PostTweet)
	})

	return h
}

func feedHandlers(rt *Router) http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Use(gatekeeper)
		r.Get("/", rt.GetFeeds)
	})

	return h
}
