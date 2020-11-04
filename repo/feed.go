package repo

import "github.com/ark1790/alpha/model"

// Feed represents Feed repository interface
type Feed interface {
	EnsureIndices(*model.Feed) error
	List(username string, t string) ([]model.Feed, error)
	Create(user *model.Feed) error
}
