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

// Mock ProductService
type mockProductService struct {
	mock.Mock
}

func (m *mockProductService) CreateProduct(product *api.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductService) GetProduct(id int) (*api.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.Product), args.Error(1)
}

func (m *mockProductService) UpdateProduct(id int, product *api.Product) error {
	args := m.Called(id, product)
	return args.Error(0)
}

func (m *mockProductService) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockProductService) ListProducts(storeID string, params map[string]interface{}) ([]api.Product, error) {
	args := m.Called(storeID, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]api.Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductService)
	controller := NewProductController(mockService)

	// Add route handler
	app.Post("/api/stores/:store_id/products", controller.CreateProduct)

	tests := []struct {
		name           string
		storeID        string
		body           api.Product
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name:    "Success",
			storeID: "store-123",
			body: api.Product{
				ProductMasterID: "pm-123",
				Price:          10000,
				StockQuantity:  100,
				Photo:           "product-1.jpg",
			},
			expectedStatus: fiber.StatusCreated,
			mockBehavior: func() {
				mockService.On("CreateProduct", mock.AnythingOfType("*api.Product")).Return(nil)
			},
		},
		{
			name:    "Service Error",
			storeID: "store-123",
			body: api.Product{
				ProductMasterID: "pm-123",
				Price:          10000,
				StockQuantity:  100,
				Photo:           "product-1.jpg",
			},
			expectedStatus: fiber.StatusInternalServerError,
			mockBehavior: func() {
				mockService.On("CreateProduct", mock.AnythingOfType("*api.Product")).Return(fmt.Errorf("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			jsonBody, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/stores/%s/products", tt.storeID), bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetProduct(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductService)
	controller := NewProductController(mockService)

	// Add route handler
	app.Get("/api/products/:id", controller.GetProduct)

	tests := []struct {
		name           string
		productID      string
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name:           "Success",
			productID:      "1",
			expectedStatus: fiber.StatusOK,
			mockBehavior: func() {
				mockService.On("GetProduct", 1).Return(&api.Product{
					ID:             1,
					ProductMasterID: "pm-123",
					Price:          10000,
					StockQuantity:  100,
					Photo:           "product-1.jpg",
				}, nil)
			},
		},
		{
			name:           "Not Found",
			productID:      "1",
			expectedStatus: fiber.StatusNotFound,
			mockBehavior: func() {
				mockService.On("GetProduct", 1).Return(nil, fmt.Errorf("not found"))
			},
		},
		{
			name:           "Invalid ID",
			productID:      "invalid",
			expectedStatus: fiber.StatusBadRequest,
			mockBehavior:   func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req := httptest.NewRequest("GET", fmt.Sprintf("/api/products/%s", tt.productID), nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}

func TestListProducts(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductService)
	controller := NewProductController(mockService)

	// Add route handler
	app.Get("/api/stores/:store_id/products", controller.ListProducts)

	tests := []struct {
		name           string
		storeID        string
		query          string
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name:           "Success",
			storeID:        "store-123",
			query:          "?page=1&limit=10",
			expectedStatus: fiber.StatusOK,
			mockBehavior: func() {
				mockService.On("ListProducts", "store-123", mock.AnythingOfType("map[string]interface{}")).
					Return([]api.Product{{
						ID:             1,
						ProductMasterID: "pm-123",
						Price:          10000,
						StockQuantity:  100,
						Photo:           "product-1.jpg",
					}}, nil)
			},
		},
		{
			name:           "Service Error",
			storeID:        "store-123",
			query:          "?page=1&limit=10",
			expectedStatus: fiber.StatusInternalServerError,
			mockBehavior: func() {
				mockService.On("ListProducts", "store-123", mock.AnythingOfType("map[string]interface{}")).
					Return(nil, fmt.Errorf("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req := httptest.NewRequest("GET", fmt.Sprintf("/api/stores/%s/products%s", tt.storeID, tt.query), nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	app := fiber.New()
	mockService := new(mockProductService)
	controller := NewProductController(mockService)

	// Add route handler
	app.Put("/api/products/:id", controller.UpdateProduct)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		mockBehavior   func()
	}{
		{
			name: "Success",
			requestBody: map[string]interface{}{
				"price":          15000,
				"stock_quantity": 50,
				"photo":          "updated-product.jpg",
			},
			expectedStatus: fiber.StatusOK,
			mockBehavior: func() {
				mockService.On("UpdateProduct", 1, mock.AnythingOfType("*api.Product")).Return(nil)
			},
		},
		{
			name:           "Service Error",
			requestBody:    map[string]interface{}{},
			expectedStatus: fiber.StatusInternalServerError,
			mockBehavior: func() {
				mockService.On("UpdateProduct", 1, mock.AnythingOfType("*api.Product")).Return(fmt.Errorf("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			jsonBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", "/api/products/1", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			mockService.AssertExpectations(t)
		})
	}
}
