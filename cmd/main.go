package main

import (
    "log"
    "net/http"
    "sync"
    "time"

    "example.com/myapp/internal/api"
    "example.com/myapp/internal/services"
    "example.com/myapp/pkg/database"
)

func main() {
    // Initialize database
    db := database.NewDB()
    
    // Initialize services
    userService := services.NewUserService(db)
    productService := services.NewProductService(db)
    
    // Start background tasks with goroutines
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        runBackgroundTasks()
    }()
    
    // Setup and start HTTP server
    router := api.SetupRoutes(userService, productService)
    
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
    
    wg.Wait()
}

func runBackgroundTasks() {
    ticker := time.NewTicker(5 * time.Minute)
    for {
        select {
        case <-ticker.C:
            log.Println("Running background task...")
        }
    }
} 