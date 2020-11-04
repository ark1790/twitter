package backend

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ark1790/alpha/api"
	"github.com/ark1790/alpha/repo"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// Server ...
type Server struct {
	userRepo   repo.User
	followRepo repo.Follow
	tweetRepo  repo.Tweet
	feedRepo   repo.Feed
}

// NewServer ...
func NewServer(ur repo.User, fl repo.Follow, tr repo.Tweet, fdr repo.Feed) *Server {
	return &Server{
		userRepo:   ur,
		followRepo: fl,
		tweetRepo:  tr,
		feedRepo:   fdr,
	}
}

// Serve ...
func (s *Server) Serve() {

	portStr := viper.GetString("PORT")

	r := chi.NewMux()
	r.Mount("/api/v1", api.NewRouter(s.userRepo, s.followRepo, s.tweetRepo, s.feedRepo))

	srvr := &http.Server{
		ReadTimeout:  viper.GetDuration("READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("WRITE_TIMEOUT"),
		IdleTimeout:  viper.GetDuration("IDLE_TIMEOUT"),
		Addr:         ":" + portStr,
		Handler:      r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Println("Server Listening on :" + portStr)
		log.Fatal(srvr.ListenAndServe())
	}()

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srvr.Shutdown(ctx)

	log.Println("Server shut down gracefully")
}
