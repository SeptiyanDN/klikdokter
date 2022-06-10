package products

type Services interface {
	SaveProduct(input ProductInput) (Product, error)
	FindAll() ([]Product, error)
	FindBySku(sku string) (Product, error)
	FindAndUpdateDataBySku(input ProductInput) (Product, error)
	RemoveDataBySku(sku string) (Product, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) SaveProduct(input ProductInput) (Product, error) {
	product := Product{}
	product.Sku = input.Sku
	product.Product_name = input.Product_name
	product.Qty = input.Qty
	product.Price = input.Price
	product.Unit = input.Unit
	product.Status = input.Status
	newProduct, err := s.repository.Save(product)
	if err != nil {
		return product, err
	}
	return newProduct, nil
}

func (s *services) FindAll() ([]Product, error) {
	Products, err := s.repository.FindAll()
	if err != nil {
		return Products, err
	}
	return Products, nil
}

func (s *services) FindBySku(sku string) (Product, error) {
	Product, err := s.repository.FindBySku(sku)
	if err != nil {
		return Product, err
	}
	return Product, nil
}

func (s *services) FindAndUpdateDataBySku(input ProductInput) (Product, error) {
	product := Product{}
	product.Sku = input.Sku
	product.Product_name = input.Product_name
	product.Qty = input.Qty
	product.Price = input.Price
	product.Unit = input.Unit
	product.Status = input.Status
	newProduct, err := s.repository.UpdateDataBySku(input.Sku, product)
	if err != nil {
		return product, err
	}
	return newProduct, nil
}

func (s *services) RemoveDataBySku(sku string) (Product, error) {
	product := Product{}
	product.Sku = sku
	newProduct, err := s.repository.RemoveDataBySku(sku)
	if err != nil {
		return product, err
	}
	return newProduct, nil
}
