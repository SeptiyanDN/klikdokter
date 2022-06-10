package products

import "gorm.io/gorm"

type Repository interface {
	Save(product Product) (Product, error)
	FindAll() ([]Product, error)
	FindBySku(sku string) (Product, error)
	UpdateDataBySku(sku string, product Product) (Product, error)
	RemoveDataBySku(sku string) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindBySku(sku string) (Product, error) {
	var Products Product
	err := r.db.Where("sku = ?", sku).Find(&Products).Error
	if err != nil {
		return Products, err
	}
	return Products, nil
}

func (r *repository) UpdateDataBySku(sku string, product Product) (Product, error) {
	err := r.db.Model(&product).Where("sku = ?", sku).Updates(product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) RemoveDataBySku(sku string) (Product, error) {
	var Products Product
	err := r.db.Where("sku = ?", sku).Delete(&Products).Error
	if err != nil {
		return Products, err
	}
	return Products, nil
}
