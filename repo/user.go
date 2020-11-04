package repo

import "github.com/ark1790/alpha/model"

// User represents user repository interface
type User interface {
	EnsureIndices(*model.User) error
	Fetch(username string) (*model.User, error)
	Create(user *model.User) error
}
