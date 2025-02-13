package services

import (
	"fmt"

	"github.com/seleraseblak/backend/api"
)

type spicyLevelService struct {
	spicyLevels []api.SpicyLevel
}

func NewSpicyLevelService() api.SpicyLevelService {
	// Data hardcode
	spicyLevels := []api.SpicyLevel{
		{ID: "1", Name: "Normal", Level: 1, Price: 0},
		{ID: "2", Name: "Pedas", Level: 2, Price: 2000},
		{ID: "3", Name: "Extra Pedas", Level: 3, Price: 4000},
		{ID: "4", Name: "Gila", Level: 4, Price: 6000},
		{ID: "5", Name: "Mati Rasa", Level: 5, Price: 8000},
	}
	return &spicyLevelService{spicyLevels: spicyLevels}
}

func (s *spicyLevelService) GetSpicyLevels() ([]api.SpicyLevel, error) {
	return s.spicyLevels, nil
}

func (s *spicyLevelService) GetSpicyLevel(id string) (*api.SpicyLevel, error) {
	for _, level := range s.spicyLevels {
		if level.ID == id {
			return &level, nil
		}
	}
	return nil, fmt.Errorf("spicy level not found")
}
