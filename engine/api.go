package engine

import (
	"database/sql"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func _databaseCheckEmailPassword(db *sql.DB, email, passwordHashed string) (bool, error) {
	err := db.QueryRow("SELECT 1 from album where email=? and hashed_password=?",
		email, passwordHashed).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("engine.login: credentials mismatched for email %s", email)
		}
		return true, fmt.Errorf("engine.login general database error")
	}

	return false, nil
}

func (e *Engine) login(email, password string) (*jwt.Token, error) {
	passwordHashed := password
	_databaseCheckEmailPassword(e.DbC, email, passwordHashed)
	return nil, nil
}
