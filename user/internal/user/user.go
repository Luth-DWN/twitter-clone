package user

import (
	"time"
)

// User represents a user
type User struct {
	ID             int64     `json:"id"`
	DisplayName    string    `json:"display_name"`
	Handle         string    `json:"handle"`
	Email          string    `json:"email"`
	PasswordDigest string    `json:"password_digest"`
	Bio            string    `json:"bio"`
	Location       string    `json:"location"`
	BirthDate      string    `json:"birth_date"`
	CreatedAt      time.Time `json:"created_at"`
}

// PasswordIsValid determines wether the given password matches with the user's password digest
func (u *User) PasswordIsValid(password Password) (bool, error) {
	return password.Compare(u.PasswordDigest)
}
