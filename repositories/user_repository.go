package repositories

import (
	"go-rest-api/config"
	"go-rest-api/models"
)

func CreateUser(user *models.User) error {
	result, err := config.DB.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
