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

type mockStoreService struct {
    mock.Mock
}

func (m *mockStoreService) CreateStore(store *api.Store) error {
    args := m.Called(store)
    return args.Error(0)
}

func (m *mockStoreService) GetStore(id string) (*api.Store, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*api.Store), args.Error(1)
}

func (m *mockStoreService) UpdateStore(id string, store *api.Store) error {
    args := m.Called(id, store)
    return args.Error(0)
}

func (m *mockStoreService) DeleteStore(id string) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *mockStoreService) ListStores(params map[string]interface{}) ([]api.Store, error) {
    args := m.Called(params)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]api.Store), args.Error(1)
}

func TestCreateStore(t *testing.T) {
    app := fiber.New()
    mockService := new(mockStoreService)
    controller := NewStoreController(mockService)
    app.Post("/api/stores", controller.CreateStore)

    tests := []struct {
        name           string
        body           api.Store
        expectedStatus int
        mockBehavior   func()
    }{
        {
            name: "Success",
            body: api.Store{
                StoreName:    "Test Store",
                StoreAddress: "Test Address",
                StorePhone:   "1234567890",
            },
            expectedStatus: fiber.StatusCreated,
            mockBehavior: func() {
                mockService.On("CreateStore", mock.AnythingOfType("*api.Store")).Return(nil)
            },
        },
        {
            name: "Service Error",
            body: api.Store{
                StoreName:    "Test Store",
                StoreAddress: "Test Address",
                StorePhone:   "1234567890",
            },
            expectedStatus: fiber.StatusInternalServerError,
            mockBehavior: func() {
                mockService.On("CreateStore", mock.AnythingOfType("*api.Store")).Return(fmt.Errorf("service error"))
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.mockBehavior()

            jsonBody, _ := json.Marshal(tt.body)
            req := httptest.NewRequest("POST", "/api/stores", bytes.NewBuffer(jsonBody))
            req.Header.Set("Content-Type", "application/json")

            resp, err := app.Test(req)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedStatus, resp.StatusCode)

            mockService.AssertExpectations(t)
        })
    }
}

func TestGetStore(t *testing.T) {
    app := fiber.New()
    mockService := new(mockStoreService)
    controller := NewStoreController(mockService)
    app.Get("/api/stores/:id", controller.GetStore)

    tests := []struct {
        name           string
        storeID       string
        expectedStatus int
        mockBehavior   func()
    }{
        {
            name:           "Success",
            storeID:       "store-123",
            expectedStatus: fiber.StatusOK,
            mockBehavior: func() {
                mockService.On("GetStore", "store-123").Return(&api.Store{
                    ID:           "store-123",
                    StoreName:    "Test Store",
                    StoreAddress: "Test Address",
                    StorePhone:   "1234567890",
                }, nil)
            },
        },
        {
            name:           "Not Found",
            storeID:       "store-123",
            expectedStatus: fiber.StatusNotFound,
            mockBehavior: func() {
                mockService.On("GetStore", "store-123").Return(nil, fmt.Errorf("not found"))
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.mockBehavior()

            req := httptest.NewRequest("GET", fmt.Sprintf("/api/stores/%s", tt.storeID), nil)
            resp, err := app.Test(req)

            assert.NoError(t, err)
            assert.Equal(t, tt.expectedStatus, resp.StatusCode)

            mockService.AssertExpectations(t)
        })
    }
}
