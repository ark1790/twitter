package repo

import "github.com/ark1790/alpha/model"

// Tweet represents user repository interface
type Tweet interface {
	EnsureIndices(*model.Tweet) error
	List(username string) (*model.Tweet, error)
	Create(user *model.User) error
}
