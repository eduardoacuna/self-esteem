package users

import (
	"context"

	"github.com/eduardoacuna/self-esteem/database"
	"github.com/eduardoacuna/self-esteem/log"
	"github.com/eduardoacuna/self-esteem/notifications"
)

// User represents a self-esteem user
type User struct {
	ID     int                   `json:"id"`
	Email  string                `json:"email"`
	Notify notifications.Rate    `json:"notify"`
	Day    notifications.Weekday `json:"day"`
}

// GetAllUsers returns all self-esteem users
func GetAllUsers(ctx context.Context) ([]*User, error) {
	log.Debug(ctx, "get all users")
	conn := database.Connection()
	q := "SELECT * FROM users"
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Notify, &user.Day); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, rows.Err()
	}

	return users, nil
}

// GetUserByID returns the user with the given ID
func GetUserByID(ctx context.Context, userID int) (*User, error) {
	log.Debug(ctx, "get user by id", "userID", userID)
	conn := database.Connection()
	q := "SELECT * FROM users WHERE id=$1"
	row := conn.QueryRow(q, userID)
	user := &User{}
	err := row.Scan(&user.ID, &user.Email, &user.Notify, &user.Day)
	if err != nil {
		return nil, err
	}
	return user, nil
}
