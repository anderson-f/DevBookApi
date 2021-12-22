package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository create a users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user on database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, username, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Search get all users that attends a filter by name or username
func (repository Users) Search(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // %nameOrNick%

	lines, err := repository.db.Query(
		"select id, name, username, email, createdAt from users where name like ? or username like ?",
		nameOrUsername, nameOrUsername,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchByID bring a user from database
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	line, err := repository.db.Query(
		"select id, name, username, email, createdAt from users where id = ?",
		ID,
	)

	if err != nil {
		return models.User{}, err
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update change the informations from users on database
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, username = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete erases an user on database
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// SearchByEmail bring a user from database by email and return the id an password
func (repository Users) SearchByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err := line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
