package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Router ...
type Router struct {
	*chi.Mux
}

// NewRouter ...
func NewRouter() *Router {
	router := &Router{
		Mux: chi.NewRouter(),
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
}
