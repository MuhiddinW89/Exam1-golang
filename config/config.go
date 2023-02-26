package config

type Config struct {
	UserFileName       string
	ProductFileName    string
	ShopCartFileName   string
	CommissionFileName string
	CategoryName       string
	OrderFileName      string
	OrderItemsFileName string
	BranchesFileName   string
}

func Load() Config {
	cfg := Config{}

	cfg.UserFileName = "./data/user.json"
	cfg.ProductFileName = "./data/product.json"
	cfg.ShopCartFileName = "./data/shop_cart.json"
	cfg.CommissionFileName = "./data/commission.json"
	cfg.CategoryName = "./data/category.json"
	cfg.OrderFileName = "./data/order.json"
	cfg.OrderItemsFileName = "./data/orderItems.json"
	cfg.BranchesFileName = "./data/branch.json"

	return cfg
}
