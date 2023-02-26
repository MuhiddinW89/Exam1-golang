package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)


	shopCartDataFilter(c) 			// 1 and 12 task
	// clientHistory(c) 				// 2 task
	// clientTotalMoneySpend(c) 		// 3 task
	// productSoldCoun(c) 				// 4 task
	// top10ProductSoldCount(c) 		// 5 task
	// top10LowSoldProductCount(c) 		// 6 task
	// mostProductSoldDate(c) 				// 7 task chala
	// productSoldByCatygories(c) 		// 8 task
	// mostActiveClient(c) 				// 9 task
	// discountToClient(c) 				// 10 task
	// createBranch(c) 					// 11 task

}

func shopCartDataFilter(c *controller.Controller){
	data, err := c.ShopCartDataFilter("2022-06-25","2022-06-25")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func clientHistory(c *controller.Controller) {
	c.ClientHistory("Misty")
}

func clientTotalMoneySpend(c *controller.Controller){
	c.ClientTotalMoneySpend("Misty")
}

func productSoldCoun(c *controller.Controller){
	c.ProductSoldCount("Qora futbolka")
}

func top10ProductSoldCount(c *controller.Controller){
	c.Top10ProductSoldCount()
}

func top10LowSoldProductCount (c *controller.Controller) {
	c.Top10LowSoldProductCount()
}

func mostProductSoldDate(c *controller.Controller){
	c.MostProductSoldDate()
}

func productSoldByCatygories(c *controller.Controller){
	c.ProductSoldByCatygories()
}

func mostActiveClient(c *controller.Controller){
	c.MostActiveClient()
}

func discountToClient(c *controller.Controller){
	c.DiscountToClient("Sherali")
}

func createBranch(c *controller.Controller) {
	c.CreateBranch(models.CreateBranch{Name: "Toshkent"})
}



// func createShopcart (c *controller.Controller) {
// 	c.AddShopCart( &models.Add{
// 		ProductId: "a80cc924-fec3-4717-8289-f23604de45ae", 
// 		UserId: "ebea6d88-820e-4863-8f69-e91f891b9000",
// 		Count: 4,
// 	})
// }


//a80cc924-fec3-4717-8289-f23604de45ae
//ac307233-33b8-425a-9d4d-fb2ddc6b4c21







// func Order(c *controller.Controller) {
	// id,err := c.CreateOrder(models.CreateOrder{
	// 	Cutomer_name: "Anvar",
	// 	Customer_address: "123 Main St.",
	// 	Customer_phone: "123-456-7890",
	// 	OrderItems: []models.CreateOrderItems{
	// 		{Product_id: "31216468-60bd-4694-b5a8-6da80febfdf6",
	// 		Count: 5},
	// 	},
	// })
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(id)

// 	res ,err := c.GetByIdOrder("7ffb629e-de2e-4842-a488-82b4588e0254")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(res)
// }

// func Product(c *controller.Controller) {

// 	c.CreateProduct(&models.CreateProduct{
// 		Name:       "Smartfon vivo V25 8/256 GB",
// 		Price:      4_860_000,
// 		CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// 	})

// 	product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	products, err := c.GetAllProduct(
// 		&models.GetListProductRequest{
// 			Offset:     0,
// 			Limit:      1,
// 			CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// 		},
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for in, product := range products.Products {
// 		fmt.Println(in+1, product)
// 	}
// }

// func Category(c *controller.Controller) {
// 	c.CreateCategory(&models.CreateCategory{
// 		Name:     "Smartfonlar va telefonlar",
// 		ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
// 	})

// 	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(category)

// }

// func User(c *controller.Controller) {

// 	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
// 	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
// 	err := c.MoneyTransfer(sender, receiver, 500_000)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }




/*
[
  {
    "id": "2d0a65a6-8a5d-4850-9211-38a78cc37595",
    "productId": "ce06cf2e-6577-46cb-96a3-1cea379bde4b",
    "userID": "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
    "count": 5,
    "status": true
  },
  {
    "id": "13675370-c9c3-4421-b8db-4841479b6a8c",
    "productId": "4e14caca-9456-4951-8934-e3fd17b4305a",
    "userID": "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
    "count": 1,
    "status": true
  },
  {
    "id": "b4fa29d4-36a7-4e5b-b739-9136387bbddf",
    "productId": "4f148546-fa00-4f6c-a989-7f9fb90faf99",
    "userID": "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
    "count": 4,
    "status": true
  }
]
*/
