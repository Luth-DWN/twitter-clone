package backend

import "github.com/HotPotatoC/twitter-clone/user/internal/user"

type User interface {
	Create(handle, email, hash string) error
	Delete(id string) error
	Update(id, displayName, bio, location, website, birthDate string) error
	GetByID(id string) (user.User, error)
}
