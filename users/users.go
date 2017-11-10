package users

import (
	"fmt"

	"github.com/eduardoacuna/self-esteem/notifications"
)

// User represents a self-esteem user
type User struct {
	ID     string                `json:"id"`
	Email  string                `json:"email"`
	Notify notifications.Rate    `json:"notify"`
	Day    notifications.Weekday `json:"day"`
}

// GetAllUsers returns all self-esteem users
func GetAllUsers() ([]*User, error) {
	return dummyUserData, nil
}

// GetUserByID returns the user with the given ID
func GetUserByID(userID string) (*User, error) {
	for _, user := range dummyUserData {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with id %v not found", userID)
}
