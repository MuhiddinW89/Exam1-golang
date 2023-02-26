package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Commission() CommissionRepoI
	Category() CategoryRepoI
	Order() OrderRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Read() ([]models.User, error)
	Create(*models.CreateUser) (string, error)
	Delete(*models.UserPrimaryKey) error
	Update(*models.UpdateUser, string) error
	GetByID(*models.UserPrimaryKey) (models.User, error)
	GetAll(*models.GetListRequest) (models.GetListResponse, error)
}

type ProductRepoI interface {
	Read() ([]models.Product, error)
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (models.ProductWithCategory, error)
	GetAll(*models.GetListProductRequest) (models.GetListProduct, error)
	Update(*models.UpdateProduct, string) error
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartRepoI interface {
	Read() ([]models.ShopCart, error)
	AddShopCart(*models.Add) (string, error)
	RemoveShopCart(*models.Remove) error
	GetUserShopCart(*models.UserPrimaryKey) ([]models.ShopCart, error)
	UpdateShopCart(string) error
}

type CommissionRepoI interface {
	AddCommission(*models.Commission) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}

type OrderRepoI interface {
	CreateOrder(order models.CreateOrder,total int) (string, error)
	GetByIdOrder(id string) (models.GetByIdOrder, error)	
}

type BranchRepoI interface{
	Read() ([]models.Branch, error)
	CreateBranch(req models.CreateBranch) error
	DeleteBranch(req *models.BranchPrimaryKey) error
	UpdateBranch(req *models.BranchPrimaryKey, branchId string) error
	GetByID(req *models.UserPrimaryKey) (models.Branch, error)
}