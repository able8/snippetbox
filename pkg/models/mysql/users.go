package mysql

import (
	"database/sql"

	"github.com/able8/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Verify whether a user exists with the provided email address and password.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get details for a specific user based on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
