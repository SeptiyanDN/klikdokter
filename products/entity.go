package products

import "time"

type Product struct {
	ID           int
	Sku          string
	Product_name string
	Qty          int64
	Price        float64
	Unit         string
	Status       int
	Created_at   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Updated_at   time.Time
}
