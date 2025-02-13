package services

import (
	"fmt"

	"github.com/seleraseblak/backend/api"
	"gorm.io/gorm"
)

type productMasterService struct {
	db *gorm.DB
}

func NewProductMasterService(db *gorm.DB) api.ProductMasterService {
	return &productMasterService{db: db}
}

func (s *productMasterService) CreateProductMaster(product *api.ProductMaster) error {
	// Validasi kategori
	if product.Category == nil {
		product.Category = []string{} // Pastikan tidak nil
	}

	product.Status = api.StatusDraft
	return s.db.Create(product).Error
}

func (s *productMasterService) GetProductMaster(id string) (*api.ProductMaster, error) {
	var product api.ProductMaster
	err := s.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *productMasterService) UpdateProductMaster(id string, product *api.ProductMaster) error {
	return s.db.Model(&api.ProductMaster{}).Where("id = ?", id).Updates(product).Error
}

func (s *productMasterService) DeleteProductMaster(id string) error {
	return s.db.Model(&api.ProductMaster{}).Where("id = ?", id).Update("status", api.StatusArchived).Error
}

func (s *productMasterService) ListProductMasters(params map[string]interface{}) ([]api.ProductMaster, error) {
	var products []api.ProductMaster
	query := s.db.Model(&api.ProductMaster{}).Where("status = ?", api.StatusPublished)

	// Apply search filter
	if search, ok := params["search"].(string); ok && search != "" {
		query = query.Where("product_name ILIKE ?", "%"+search+"%")
	}

	// Apply category filter jika ada
	if category, ok := params["category"].(string); ok && category != "" {
		// Gunakan operator @> untuk mencari dalam array JSON
		query = query.Where("category @> ?", fmt.Sprintf("[\"%s\"]", category))
	}

	// Apply pagination
	page := 1
	limit := 10
	if p, ok := params["page"].(int); ok {
		page = p
	}
	if l, ok := params["limit"].(int); ok {
		limit = l
	}
	offset := (page - 1) * limit

	err := query.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
