package models

import (
	"errors"
	"strings"
	"time"

	"api/src/security"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if user.Username == "" {
		return errors.New("O Username é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail inserido é invalido")
	}

	if stage == "create" && user.Password == "" {
		return errors.New("O senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "create" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}
