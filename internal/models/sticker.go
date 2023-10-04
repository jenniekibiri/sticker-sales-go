package models

import (
	"gorm.io/gorm"
)

type Sticker struct {
	gorm.Model
	Title       string  `gorm:"unique" json:"title"`
	Description string  `json:"description" `
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Images      string  `json:"images"`
	Size        string  `gorm:"check: size IN ('3X3', '4X4', '5X5', '6X6', '7X7', '8X8', '9X9', '10X10')" json:"size"`
}

type Order struct {
	gorm.Model
	Stickers        string  `json:"stickers"`
	Total           float64 `json:"total"`
	PaymentMethodId string  `json:"payment_method_id"`
}
