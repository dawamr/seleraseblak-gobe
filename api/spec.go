package api

import (
	"time"
)

// Request/Response structures
type Store struct {
	ID           string `json:"id" gorm:"primaryKey;type:uuid;column:id"`
	StoreName    string `json:"store_name" gorm:"column:store_name"`
	StoreAddress string `json:"store_address" gorm:"column:store_address"`
	StorePhone   string `json:"store_phone" gorm:"column:store_phone"`
	Status       string `json:"status" gorm:"column:status"`
	DateCreated  string `json:"date_created" gorm:"column:date_created"`
	DateUpdated  string `json:"date_updated" gorm:"column:date_updated"`
}

func (Store) TableName() string {
	return "Store"
}

type ProductMaster struct {
	ID          string   `json:"id" gorm:"primaryKey;type:uuid;column:id"`
	ProductName string   `json:"product_name" gorm:"column:product_name"`
	Category    []string `json:"category" gorm:"serializer:json;column:category"`
	SKU         string   `json:"sku" gorm:"column:sku"`
	Description string   `json:"description" gorm:"column:description"`
	Status      string   `json:"status" gorm:"column:status"`
	UserCreated string   `json:"user_created" gorm:"column:user_created;type:uuid"`
	UserUpdated string   `json:"user_updated" gorm:"column:user_updated;type:uuid"`
	Price       int      `json:"price" gorm:"column:price"`
}

func (ProductMaster) TableName() string {
	return "Product_Master"
}

type Product struct {
	ID              int              `json:"id" gorm:"primaryKey;column:id"`
	ProductMasterID string           `json:"product_master_id" gorm:"column:product_master_id;type:uuid"`
	ProductMaster   ProductMaster    `json:"product_master" gorm:"foreignKey:ProductMasterID"`
	StoreID         string           `json:"store_id" gorm:"column:store_id;type:uuid"`
	Price           float64          `json:"price" gorm:"column:price"`
	StockQuantity   int              `json:"stock_quantity" gorm:"column:stock_quantity"`
	IsActive        bool             `json:"is_active" gorm:"column:is_active"`
	Status          string           `json:"status" gorm:"column:status"`
	Photo           string           `json:"photo" gorm:"column:photo"`
	ProductToppings []ProductTopping `json:"-" gorm:"foreignKey:Product_id"`
	Toppings        []Topping        `json:"toppings" gorm:"-"`
}

func (Product) TableName() string {
	return "Product"
}

// Tambahkan method untuk mengkonversi ProductToppings ke Toppings
func (p *Product) AfterFind() error {
	if len(p.ProductToppings) > 0 {
		p.Toppings = make([]Topping, 0)
		for _, pt := range p.ProductToppings {
			if pt.Topping.ID != 0 { // Pastikan topping valid
				p.Toppings = append(p.Toppings, pt.Topping)
			}
		}
	}
	return nil
}

// Tambahkan tabel junction
type ProductTopping struct {
	ID        int     `gorm:"primaryKey;column:id"`
	ProductID int     `gorm:"column:Product_id"`
	ToppingID int     `gorm:"column:Topping_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:id"`
	Topping   Topping `gorm:"foreignKey:ToppingID;references:id"`
}

func (ProductTopping) TableName() string {
	return "Product_Topping"
}

type UserStore struct {
	ID          int    `json:"id" gorm:"primaryKey;column:id"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:uuid"`
	StoreID     string `json:"store_id" gorm:"column:store_id;type:uuid"`
	RoleInStore string `json:"role_in_store" gorm:"column:role_in_store"`
	Status      string `json:"status" gorm:"column:status"`
}

func (UserStore) TableName() string {
	return "User_Store"
}

// Tambahkan konstanta untuk status
const (
	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusArchived  = "archived"
)

// Tambahkan struct Topping
type Topping struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id"`
	Status      string    `json:"status" gorm:"column:status"`
	DateCreated time.Time `json:"date_created" gorm:"column:date_created"`
	DateUpdated time.Time `json:"date_updated" gorm:"column:date_updated"`
	Price       int       `json:"price" gorm:"column:price"`
	Name        string    `json:"name" gorm:"column:name"`
}

func (Topping) TableName() string {
	return "Topping"
}

// Tambahkan struct SpicyLevel
type SpicyLevel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Price int    `json:"price"`
}

// Tambahkan interface service
type SpicyLevelService interface {
	GetSpicyLevels() ([]SpicyLevel, error)
	GetSpicyLevel(id string) (*SpicyLevel, error)
}

// Service interfaces
type StoreService interface {
	CreateStore(store *Store) error
	GetStore(id string) (*Store, error)
	UpdateStore(id string, store *Store) error
	DeleteStore(id string) error
	ListStores(params map[string]interface{}) ([]Store, error)
}

type ProductService interface {
	CreateProduct(product *Product) error
	GetProduct(id int) (*Product, error)
	UpdateProduct(id int, product *Product) error
	DeleteProduct(id int) error
	ListProducts(storeID string, params map[string]interface{}) ([]Product, error)
}

type ProductMasterService interface {
	CreateProductMaster(product *ProductMaster) error
	GetProductMaster(id string) (*ProductMaster, error)
	UpdateProductMaster(id string, product *ProductMaster) error
	DeleteProductMaster(id string) error
	ListProductMasters(params map[string]interface{}) ([]ProductMaster, error)
}

type UserStoreService interface {
	AssignUserToStore(userStore *UserStore) error
	RemoveUserFromStore(userID, storeID string) error
	GetUserStores(userID string) ([]UserStore, error)
	GetStoreUsers(storeID string) ([]UserStore, error)
}

type ToppingService interface {
	GetToppings() ([]Topping, error)
	GetTopping(id int) (*Topping, error)
	CreateTopping(topping *Topping) error
	UpdateTopping(id int, topping *Topping) error
	DeleteTopping(id int) error
}

// Tambahkan interface service
type ProductToppingService interface {
	GetProductToppings() ([]ProductTopping, error)
	GetProductToppingsByProduct(productID int) ([]ProductTopping, error)
	GetProductToppingsByTopping(toppingID int) ([]ProductTopping, error)
	CreateProductTopping(productTopping *ProductTopping) error
	DeleteProductTopping(productID, toppingID int) error
}
