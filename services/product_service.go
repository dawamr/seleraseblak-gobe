package services

import (
	"github.com/seleraseblak/backend/api"
	"gorm.io/gorm"
)

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) api.ProductService {
	return &productService{db: db}
}

func (s *productService) CreateProduct(product *api.Product) error {
	product.Status = api.StatusDraft
	if product.Photo != "" {
		// Tambahkan validasi format/ukuran photo jika diperlukan
	}
	return s.db.Create(product).Error
}

func (s *productService) GetProduct(id int) (*api.Product, error) {
	var product api.Product
	err := s.db.Preload("ProductMaster").
		Preload("ProductToppings.Topping", "status = ?", api.StatusPublished).
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *productService) UpdateProduct(id int, product *api.Product) error {
	if product.Photo != "" {
		// Tambahkan validasi format/ukuran photo jika diperlukan
	}
	return s.db.Model(&api.Product{}).Where("id = ?", id).Updates(product).Error
}

func (s *productService) DeleteProduct(id int) error {
	return s.db.Model(&api.Product{}).Where("id = ?", id).Update("status", api.StatusArchived).Error
}

func (s *productService) ListProducts(storeID string, params map[string]interface{}) ([]api.Product, error) {
	var products []api.Product
	query := s.db.Model(&api.Product{}).
		Where("store_id = ? AND status = ?", storeID, api.StatusPublished).
		Preload("ProductMaster").
		Preload("ProductToppings.Topping", "status = ?", api.StatusPublished).
		Limit(20)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	for i := range products {
		if err := products[i].AfterFind(); err != nil {
			return nil, err
		}
	}

	return products, nil
}
