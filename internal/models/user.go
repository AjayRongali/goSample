package models

import "time"

type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Example of method on struct
func (u *User) IsValid() bool {
    return u.Username != "" && u.Email != ""
}

// Example of interface
type UserRepository interface {
    Create(user *User) error
    GetByID(id int64) (*User, error)
    Update(user *User) error
    Delete(id int64) error
} 