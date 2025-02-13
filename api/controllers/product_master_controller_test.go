package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock ProductMasterService
type mockProductMasterService struct {
	mock.Mock
}

func (m *mockProductMasterService) CreateProductMaster(product *api.ProductMaster) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductMasterService) GetProductMaster(id string) (*api.ProductMaster, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.ProductMaster), args.Error(1)
}

func (m *mockProductMasterService) UpdateProductMaster(id string, product *api.ProductMaster) error {
	args := m.Called(id, product)
	return args.Error(0)
}

func (m *mockProductMasterService) DeleteProductMaster(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockProductMasterService) ListProductMasters(params map[string]interface{}) ([]api.ProductMaster, error) {
	args := m.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]api.ProductMaster), args.Error(1)
}

func TestCreateProductMaster(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductMasterService)
	controller := NewProductMasterController(mockService)

	app.Post("/api/product-masters", controller.CreateProductMaster)

	tests := []struct {
		name           string
		body           api.ProductMaster
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name: "Success",
			body: api.ProductMaster{
				ProductName: "Test Product",
				Category:    []string{"makanan", "minuman"},
				Price:      10000,
				Status:     api.StatusDraft,
			},
			expectedStatus: fiber.StatusCreated,
			mockBehavior: func() {
				mockService.On("CreateProductMaster", mock.AnythingOfType("*api.ProductMaster")).Return(nil)
			},
		},
		{
			name: "Service Error",
			body: api.ProductMaster{
				ProductName: "Test Product",
				Category:    []string{"makanan"},
				Price:      10000,
			},
			expectedStatus: fiber.StatusInternalServerError,
			mockBehavior: func() {
				mockService.On("CreateProductMaster", mock.AnythingOfType("*api.ProductMaster")).Return(fmt.Errorf("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			jsonBody, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/api/product-masters", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetProductMaster(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductMasterService)
	controller := NewProductMasterController(mockService)

	app.Get("/api/product-masters/:id", controller.GetProductMaster)

	tests := []struct {
		name           string
		productID      string
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name:           "Success",
			productID:      "123e4567-e89b-12d3-a456-426614174000",
			expectedStatus: fiber.StatusOK,
			mockBehavior: func() {
				mockService.On("GetProductMaster", "123e4567-e89b-12d3-a456-426614174000").Return(&api.ProductMaster{
					ID:          "123e4567-e89b-12d3-a456-426614174000",
					ProductName: "Test Product",
					Category:    []string{"makanan"},
					Price:      10000,
				}, nil)
			},
		},
		{
			name:           "Not Found",
			productID:      "123e4567-e89b-12d3-a456-426614174000",
			expectedStatus: fiber.StatusNotFound,
			mockBehavior: func() {
				mockService.On("GetProductMaster", "123e4567-e89b-12d3-a456-426614174000").Return(nil, fmt.Errorf("not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req := httptest.NewRequest("GET", fmt.Sprintf("/api/product-masters/%s", tt.productID), nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}

func TestListProductMasters(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductMasterService)
	controller := NewProductMasterController(mockService)

	app.Get("/api/product-masters", controller.ListProductMasters)

	tests := []struct {
		name           string
		query          string
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name:           "Success",
			query:          "?page=1&limit=10&search=test&category=makanan",
			expectedStatus: fiber.StatusOK,
			mockBehavior: func() {
				expectedParams := map[string]interface{}{
					"page":     1,
					"limit":    10,
					"search":   "test",
					"category": "makanan",
				}
				mockService.On("ListProductMasters", mock.MatchedBy(func(params map[string]interface{}) bool {
					// Verifikasi parameter yang dikirim ke service
					return params["page"] == expectedParams["page"] &&
						params["limit"] == expectedParams["limit"] &&
						params["search"] == expectedParams["search"] &&
						params["category"] == expectedParams["category"]
				})).Return([]api.ProductMaster{
					{
						ID:          "123e4567-e89b-12d3-a456-426614174000",
						ProductName: "Test Product",
						Category:    []string{"makanan"},
						Price:      10000,
					},
				}, nil)
			},
		},
		{
			name:           "Service Error",
			query:          "?page=1&limit=10",
			expectedStatus: fiber.StatusInternalServerError,
			mockBehavior: func() {
				mockService.On("ListProductMasters", mock.AnythingOfType("map[string]interface{}")).
					Return(nil, fmt.Errorf("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req := httptest.NewRequest("GET", fmt.Sprintf("/api/product-masters%s", tt.query), nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}
