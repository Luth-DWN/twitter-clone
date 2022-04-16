package user

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	// MinPasswordLength is the minimum password length required
	MinPasswordLength = 8
	// MaxPasswordLength is the maximum password length allowed
	MaxPasswordLength = 64
	// HashCostFactor is the cost factor for bcrypt
	HashCostFactor = 10
)

var (
	// ErrPasswordInvalidLength is returned when the password is not within the
	// required length range
	ErrPasswordInvalidLength = errors.New("err password invalid length")
	// ErrPasswordInvalidCharacters is returned when the password contains invalid characters
	ErrPasswordInvalidCharacters = errors.New("err password invalid characters")
)

// Password represents a password
type Password string

// Validate validates the password
func (p Password) Validate() error {
	if !p.validateLength() {
		return ErrPasswordInvalidLength
	}

	if !p.validateCharacters() {
		return ErrPasswordInvalidCharacters
	}

	return nil
}

// validateLength validates the password length
func (p Password) validateLength() bool {
	return len(p) < MinPasswordLength || len(p) > MaxPasswordLength
}

// validateCharacters validates the password characters
func (p Password) validateCharacters() bool {
	for _, char := range p {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// GenerateHash generates a hash for the password
func (p Password) GenerateHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), HashCostFactor)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare compares the password with the hash
func (p Password) Compare(digest string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(p))
	if err != nil {
		return false, err
	}

	return true, nil
}
