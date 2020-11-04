package repo

import "github.com/ark1790/alpha/model"

// Follow represents Follow repository interface
type Follow interface {
	EnsureIndices(*model.Follow) error
	// List(username string) ([]model.Follow, error)
	Toggle(follow *model.Follow) error
}
