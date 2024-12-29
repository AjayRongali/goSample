package services

import (
    "context"
    "errors"

    "example.com/myapp/internal/models"
    "example.com/myapp/pkg/database"
)

type UserService struct {
    db *database.DB
}

func NewUserService(db *database.DB) *UserService {
    return &UserService{db: db}
}

// Example of context usage and error handling
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
    if !user.IsValid() {
        return errors.New("invalid user data")
    }

    // Example of channel usage for async operations
    done := make(chan error)
    go func() {
        // Simulate database operation
        done <- s.db.Create(user)
    }()

    select {
    case err := <-done:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}

// Example of multiple return values
func (s *UserService) GetUser(id int64) (*models.User, error) {
    return s.db.GetUser(id)
} 