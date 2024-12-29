package database

import (
    "sync"
    "example.com/myapp/internal/models"
)

type DB struct {
    mu    sync.RWMutex
    users map[int64]*models.User
}

func NewDB() *DB {
    return &DB{
        users: make(map[int64]*models.User),
    }
}

// Example of mutex usage for thread safety
func (db *DB) Create(user *models.User) error {
    db.mu.Lock()
    defer db.mu.Unlock()
    
    // Simulate database operation
    db.users[user.ID] = user
    return nil
}

func (db *DB) GetUser(id int64) (*models.User, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()
    
    if user, ok := db.users[id]; ok {
        return user, nil
    }
    return nil, nil
} 