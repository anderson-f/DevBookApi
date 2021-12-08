package models

import (
	"errors"
	"strings"
	"time"
)

// User represents one user using social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Prepare will call the methods for validation e format on user received
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.cleanBorderSpaces()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if user.Username == "" {
		return errors.New("O Username é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if user.Password == "" {
		return errors.New("O senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (user *User) cleanBorderSpaces() {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}
