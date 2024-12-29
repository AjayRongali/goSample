package api

import (
    "encoding/json"
    "net/http"
    "strconv"

    "example.com/myapp/internal/models"
    "example.com/myapp/internal/services"
)

type Handler struct {
    userService    *services.UserService
    productService *services.ProductService
}

// Example of closure and middleware
func withJSON(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        handler(w, r)
    }
}

// Example of HTTP handler with error handling
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    print(r.Body)
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.userService.CreateUser(r.Context(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

// Example of URL parameter handling
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
    user, err := h.userService.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
} 
