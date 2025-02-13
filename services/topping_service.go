package services

import (
	"github.com/seleraseblak/backend/api"
	"gorm.io/gorm"
)

type toppingService struct {
	db *gorm.DB
}

func NewToppingService(db *gorm.DB) api.ToppingService {
	return &toppingService{db: db}
}

func (s *toppingService) GetToppings() ([]api.Topping, error) {
	var toppings []api.Topping
	if err := s.db.Find(&toppings).Error; err != nil {
		return nil, err
	}
	return toppings, nil
}

func (s *toppingService) GetTopping(id int) (*api.Topping, error) {
	var topping api.Topping
	if err := s.db.First(&topping, id).Error; err != nil {
		return nil, err
	}
	return &topping, nil
}

func (s *toppingService) CreateTopping(topping *api.Topping) error {
	return s.db.Create(topping).Error
}

func (s *toppingService) UpdateTopping(id int, topping *api.Topping) error {
	return s.db.Model(&api.Topping{}).Where("id = ?", id).Updates(topping).Error
}

func (s *toppingService) DeleteTopping(id int) error {
	return s.db.Delete(&api.Topping{}, id).Error
}
