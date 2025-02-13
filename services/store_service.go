package services

import (
	"time"

	"github.com/seleraseblak/backend/api"
	"gorm.io/gorm"
)

type storeService struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) api.StoreService {
	return &storeService{db: db}
}

// Implement api.StoreService interface methods

// Add interface implementations
func (s *storeService) CreateStore(store *api.Store) error {
	store.Status = api.StatusDraft
	return s.db.Create(store).Error
}

func (s *storeService) GetStore(id string) (*api.Store, error) {
	var store api.Store
	err := s.db.Where("id = ?", id).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (s *storeService) UpdateStore(id string, store *api.Store) error {
	store.DateUpdated = time.Now().Format(time.RFC3339)
	return s.db.Model(&api.Store{}).Where("id = ?", id).Updates(store).Error
}

func (s *storeService) DeleteStore(id string) error {
	return s.db.Model(&api.Store{}).Where("id = ?", id).Update("status", api.StatusArchived).Error
}

func (s *storeService) ListStores(params map[string]interface{}) ([]api.Store, error) {
	var stores []api.Store
	query := s.db.Model(&api.Store{})

	// Tambahkan filter status jika diperlukan
	if status, ok := params["status"].(string); ok {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
