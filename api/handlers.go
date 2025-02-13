package api

import "net/http"

// Handler interfaces
type StoreHandler interface {
    CreateStore(w http.ResponseWriter, r *http.Request)
    GetStore(w http.ResponseWriter, r *http.Request)
    UpdateStore(w http.ResponseWriter, r *http.Request)
    DeleteStore(w http.ResponseWriter, r *http.Request)
    ListStores(w http.ResponseWriter, r *http.Request)
}

type ProductHandler interface {
    CreateProduct(w http.ResponseWriter, r *http.Request)
    GetProduct(w http.ResponseWriter, r *http.Request)
    UpdateProduct(w http.ResponseWriter, r *http.Request)
    DeleteProduct(w http.ResponseWriter, r *http.Request)
    ListProducts(w http.ResponseWriter, r *http.Request)
}

type ProductMasterHandler interface {
    CreateProductMaster(w http.ResponseWriter, r *http.Request)
    GetProductMaster(w http.ResponseWriter, r *http.Request)
    UpdateProductMaster(w http.ResponseWriter, r *http.Request)
    DeleteProductMaster(w http.ResponseWriter, r *http.Request)
    ListProductMasters(w http.ResponseWriter, r *http.Request)
}

type UserStoreHandler interface {
    AssignUserToStore(w http.ResponseWriter, r *http.Request)
    RemoveUserFromStore(w http.ResponseWriter, r *http.Request)
    GetUserStores(w http.ResponseWriter, r *http.Request)
    GetStoreUsers(w http.ResponseWriter, r *http.Request)
}
