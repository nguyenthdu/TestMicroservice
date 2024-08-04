package model

type Category struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"` // Thêm trường này
}

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"categoryId"`
}
