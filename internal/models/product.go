package models

type Product struct {
    ID          int64   `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

type ProductRepository interface {
    Create(product *Product) error
    GetByID(id int64) (*Product, error)
    List() ([]Product, error)
} 