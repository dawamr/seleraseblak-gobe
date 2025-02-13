# API Specification

## Stores

- `POST /api/stores` - Create a new store
- `GET /api/stores/{id}` - Get store details
- `PUT /api/stores/{id}` - Update store details
- `DELETE /api/stores/{id}` - Delete a store
- `GET /api/stores` - List stores (with filtering/pagination)

## Products

- `POST /api/stores/{store_id}/products` - Add product to store
- `GET /api/stores/{store_id}/products/{id}` - Get product details
- `PUT /api/stores/{store_id}/products/{id}` - Update product
- `DELETE /api/stores/{store_id}/products/{id}` - Delete product
- `GET /api/stores/{store_id}/products` - List store products

## Product Masters

- `POST /api/product-masters` - Create product master
- `GET /api/product-masters/{id}` - Get product master
- `PUT /api/product-masters/{id}` - Update product master
- `DELETE /api/product-masters/{id}` - Delete product master
- `GET /api/product-masters` - List product masters

## User Store Management

- `POST /api/stores/{store_id}/users` - Assign user to store
- `DELETE /api/stores/{store_id}/users/{user_id}` - Remove user from store
- `GET /api/users/{user_id}/stores` - Get user's stores
- `GET /api/stores/{store_id}/users` - Get store's users

All endpoints should support:

- Authentication via JWT
- Proper error responses
- Pagination for list endpoints
- Filtering and sorting where applicable
