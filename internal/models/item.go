package models

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Quantity int     `json:"quantity" binding:"required,gt=0"`
	Price    float64 `json:"price" binding:"gte=0"`
	Bought   bool    `json:"bought"`
}