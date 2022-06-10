package products

type ProductInput struct {
	Sku          string  `form:"sku" binding:"required"`
	Product_name string  `form:"product_name" binding:"required"`
	Qty          int64   `form:"qty" binding:"required"`
	Price        float64 `form:"price" binding:"required"`
	Unit         string  `form:"unit" binding:"required"`
	Status       int     `form:"status" binding:"required"`
}

type SkuInput struct {
	Sku string `form:"sku" binding:"required"`
}
