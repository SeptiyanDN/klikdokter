package products

import "time"

type ProductsFormatter struct {
	ID           int       `json:"id"`
	Sku          string    `json:"sku"`
	Product_name string    `json:"product_name"`
	Qty          int64     `json:"qty"`
	Price        float64   `json:"price"`
	Unit         string    `json:"unit"`
	Status       int       `json:"status"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

func FormatProduct(Product Product) ProductsFormatter {
	ProductsFormatter := ProductsFormatter{}
	ProductsFormatter.ID = Product.ID
	ProductsFormatter.Sku = Product.Sku
	ProductsFormatter.Product_name = Product.Product_name
	ProductsFormatter.Qty = Product.Qty
	ProductsFormatter.Price = Product.Price
	ProductsFormatter.Unit = Product.Unit
	ProductsFormatter.Status = Product.Status
	ProductsFormatter.Created_at = Product.Created_at
	ProductsFormatter.Updated_at = Product.Updated_at

	return ProductsFormatter
}

func FormatProducts(Products []Product) []ProductsFormatter {
	ProductsFormatter := []ProductsFormatter{}
	for _, Product := range Products {
		ProductFormatter := FormatProduct(Product)
		ProductsFormatter = append(ProductsFormatter, ProductFormatter)
	}
	return ProductsFormatter
}
