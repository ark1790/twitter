package api

import (
	"log"
	"net/http"

	"github.com/ark1790/alpha/model"
)

type postTweetPld struct {
	Body string `json:"body"`
}

func (c *postTweetPld) validate() *validationError {
	errV := validationError{}

	bLen := len(c.Body)

	if bLen < 1 || bLen > 120 {
		errV.add("body", "length must be between 1 to 120")
	}

	if len(errV) > 0 {
		return &errV
	}

	return nil
}

func (rt *Router) PostTweet(w http.ResponseWriter, r *http.Request) {
	usr := getAuthUser(r)

	body := postTweetPld{}
	if err := parseBody(r, &body); err != nil {
		panic(newAPIError("Unable to parse body", errBadRequest, err))
	}
	if err := body.validate(); err != nil {
		panic(newAPIError("Invalid data", errInvalidData, err))
	}

	twt := &model.Tweet{
		Body:     body.Body,
		Username: usr,
	}

	if err := rt.tweetRepo.Create(twt); err != nil {
		panic(newAPIError("Internal Server Error", errInternalServer, err))
	}

	go rt.createFeed(twt)

	resp := response{
		code: http.StatusOK,
		Data: twt,
	}

	resp.serveJSON(w)

}

func (rt *Router) createFeed(tw *model.Tweet) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	rs, err := rt.followRepo.List(tw.Username)
	if err != nil {
		panic(err)
	}

	fd := &model.Feed{
		Body:     tw.Body,
		Username: tw.Username,
	}
	for _, fl := range rs {
		fd.For = fl.Profile
		if err := rt.feedRepo.Create(fd); err != nil {
			log.Println(err)
		}
	}

}
