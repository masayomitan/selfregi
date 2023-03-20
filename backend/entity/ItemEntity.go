package entity

import (
	// "time"
)

type ItemEntity struct {
		ID int `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		CategoryID int `json:"category_id,omitempty"`
		IsDisplay bool `json:"is_display,omitempty"`
    Tax int `json:"tax,omitempty"`
    TaxRate int `json:"tax_rate,omitempty"`
    Price int `json:"price,omitempty"`
    TemporaryStock int `json:"temporary_stock,omitempty"`
}