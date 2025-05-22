package utils

import (
	"context"
	db "musique/database"
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {

	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func UserExists(email, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM users 
			WHERE email = $1 OR user_name = $2
		)
	`
	err := db.DB.QueryRow(ctx, query, email, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
