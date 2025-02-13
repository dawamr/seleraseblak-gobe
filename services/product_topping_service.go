package services

import (
	"github.com/seleraseblak/backend/api"
	"gorm.io/gorm"
)

type productToppingService struct {
    db *gorm.DB
}

func NewProductToppingService(db *gorm.DB) api.ProductToppingService {
    return &productToppingService{db: db}
}

func (s *productToppingService) GetProductToppings() ([]api.ProductTopping, error) {
    var productToppings []api.ProductTopping
    err := s.db.Preload("Product").Preload("Topping").Find(&productToppings).Error
    return productToppings, err
}

func (s *productToppingService) GetProductToppingsByProduct(productID int) ([]api.ProductTopping, error) {
    var productToppings []api.ProductTopping
    err := s.db.Where("Product_id = ?", productID).
        Preload("Product").
        Preload("Topping").
        Find(&productToppings).Error
    return productToppings, err
}

func (s *productToppingService) GetProductToppingsByTopping(toppingID int) ([]api.ProductTopping, error) {
    var productToppings []api.ProductTopping
    err := s.db.Where("Topping_id = ?", toppingID).
        Preload("Product").
        Preload("Topping").
        Find(&productToppings).Error
    return productToppings, err
}

func (s *productToppingService) CreateProductTopping(productTopping *api.ProductTopping) error {
    return s.db.Create(productTopping).Error
}

func (s *productToppingService) DeleteProductTopping(productID, toppingID int) error {
    return s.db.Where("Product_id = ? AND Topping_id = ?", productID, toppingID).
        Delete(&api.ProductTopping{}).Error
}
