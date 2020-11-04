package repo

import "github.com/ark1790/alpha/model"

// Tweet represents tweet repository interface
type Tweet interface {
	EnsureIndices(*model.Tweet) error
	List(fUsername string) ([]model.Tweet, error)
	Create(user *model.Tweet) error
}
