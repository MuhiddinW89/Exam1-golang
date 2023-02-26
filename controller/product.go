package controller

import (
	"app/models"
	"fmt"
	"sort"
	"time"
)

func (c *Controller) CreateProduct(req *models.CreateProduct) (string, error) {
	id, err := c.store.Product().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) error {
	err := c.store.Product().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct, productId string) error {
	err := c.store.Product().Update(req, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	product, err := c.store.Product().GetByID(req)
	if err != nil {
		return models.Product{}, err
	}

	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
	if err != nil {
		return models.Product{}, err
	}

	return models.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: category,
	}, nil
}

func (c *Controller) GetAllProduct(req *models.GetListProductRequest) (models.GetListProduct, error) {
	products, err := c.store.Product().GetAll(req)
	if err != nil {
		return models.GetListProduct{}, err
	}
	return products, nil
}



func (c *Controller) ProductSoldCount(req string){
	products, err := c.store.Product().Read()
	if err != nil {
		fmt.Println(err)
	}

	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}

	for _, productsV := range products {
		if req == productsV.Name {
			i := 0
			for _, shopCartsV := range shopCarts{
				if productsV.Id == shopCartsV.ProductId && shopCartsV.Status{
					i += shopCartsV.Count
				}
			}
			fmt.Println("Product:", productsV.Name, "Sold:", i)
		}
	}
}


func (c *Controller) Top10ProductSoldCount(){
	products, err := c.store.Product().Read()
	if err != nil {
		fmt.Println(err)
	}

	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
	slice := []int{}
	myMap := map[string]int{}
	for _, productsV := range products {
			i := 0
			for _, shopCartsV := range shopCarts{
				if productsV.Id == shopCartsV.ProductId && shopCartsV.Status{
					i += shopCartsV.Count
				}
			}
			// fmt.Println(i, productsV.Name)
			slice = append(slice, i)
			myMap[productsV.Name]=i
	}
	sort.Slice(slice, func(i, j int) bool {
         return slice[i] > slice[j]
    })
	
	slice = append(slice, -1)
	uniqueSlice := []int{}
	for i := 0; i < len(slice)-1; i++ {	
		if slice[i] != slice[i+1] {
			uniqueSlice = append(uniqueSlice, slice[i] )
		}
	}
	// fmt.Println(uniqueSlice)
	i :=0 
	for _, sliceV := range uniqueSlice{
		for j, mapV := range myMap{
				if sliceV == mapV && i<10{
					i++
					fmt.Println(i, "Name:", j, "count:", mapV)
				}
			}	
	}
}


func (c *Controller) Top10LowSoldProductCount(){
	products, err := c.store.Product().Read()
	if err != nil {
		fmt.Println(err)
	}

	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
	slice := []int{}
	myMap := map[string]int{}
	for _, productsV := range products {
			i := 0
			for _, shopCartsV := range shopCarts{
				if productsV.Id == shopCartsV.ProductId && shopCartsV.Status{
					i += shopCartsV.Count
				}
			}
			// fmt.Println(i, productsV.Name)
			slice = append(slice, i)
			myMap[productsV.Name]=i
	}
	sort.Slice(slice, func(i, j int) bool {
         return slice[i] < slice[j]
    })
	
	slice = append(slice, -1)
	uniqueSlice := []int{}
	for i := 0; i < len(slice)-1; i++ {	
		if slice[i] != slice[i+1] {
			uniqueSlice = append(uniqueSlice, slice[i] )
		}
	}
	// fmt.Println(uniqueSlice)
	i :=0 
	for _, sliceV := range uniqueSlice{
		for j, mapV := range myMap{
				if sliceV == mapV && i<10{
					i++
					fmt.Println(i, "Name:", j, "count:", mapV)
				}
			}	
	}
}


func (c *Controller) MostProductSoldDate(){ // -----------------------------
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}

	defaultTimeLayout := "2006-01-02 15:04:05"
	dateSlice := []time.Time{}


	// if fromDate == "" && toDate == "" {
	// 	for _, v := range shopCarts{
	// 		dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
	// 		  dateSlice = append(dateSlice, dataInShopCart)
	// 		  slice = append(slice, v)
	// 	  }
	// }else{
	// modFromDate := fromDate + " 00:00:00"
	// modToDate := toDate + " 23:59:59"

	// strart, err1 := time.Parse(defaultTimeLayout, modFromDate)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }
	// end, err2 := time.Parse(defaultTimeLayout, modToDate)
	// if err2 != nil {
	// 	fmt.Println(err1)
	// }

	// for _, v := range shopCarts{
	//   dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
	//   if dataInShopCart.After(strart) && dataInShopCart.Before(end) {
	// 	dateSlice = append(dateSlice, dataInShopCart)
	// 	slice = append(slice, v)
	//   }
	// }
	

	for _, v := range shopCarts{
				if v.Status {
				dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
				dateSlice = append(dateSlice, dataInShopCart)
			  }
			}

	compareDates := func(i, j int) bool {
        date1 := dateSlice [i]
        date2 := dateSlice [j]
        if date1.Year() != date2.Year() {
            return date1.Year() < date2.Year()
        }
        if date1.Month() != date2.Month() {
            return date1.Month() < date2.Month()
        }
		if date1.Day() != date2.Day() {
        	return date1.Day() < date2.Day()
		}
		if date1.Hour() != date2.Hour(){
			return date1.Hour() < date2.Hour()
		}
		if date1.Minute() != date2.Minute(){
			return date1.Minute() < date2.Hour()
		}
		return date1.Second() < date2.Second()
    }
	sort.SliceStable(dateSlice, compareDates)

    // i := 0
	sortedSlice := []models.ShopCart{}	
	for _, sortedData := range dateSlice{
		for _, v := range shopCarts{
			dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
			if sortedData == dataInShopCart {
				// i++
				// fmt.Println(i,v)
				sortedSlice = append(sortedSlice, v)
			}
	}
	}
	// fmt.Println(sortedSlice[0].Time, sortedSlice[len(sortedSlice)-1].Time)


	modFromDate := sortedSlice[0].Time
	start, err1 := time.Parse(defaultTimeLayout, modFromDate)
	if err1 != nil {
		fmt.Println(err1)
	}

	start = start.Truncate(24 * time.Hour)
	end := start.AddDate(0, 0, 1) //--

	countOfSaleByDate := 0

	for i := 0; i < len(sortedSlice); i++ {
		
	for _, v := range sortedSlice{
		  dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
		  if dataInShopCart.After(start) && dataInShopCart.Before(end) {
			countOfSaleByDate += v.Count
		  }
		}
		fmt.Println(countOfSaleByDate)
		start = start.AddDate(0, 0, 1)
		end = start.AddDate(0, 0, 1)
		countOfSaleByDate = 0
		
	}	


	
	
	




}





func (c *Controller) ProductSoldByCatygories(){
	
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}

	myMap := map[string]int{}

	 i := 0
	for _, shopCartV := range shopCarts{
		if shopCartV.Status {
			i++
			count := shopCartV.Count
			product, _ := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopCartV.ProductId})
			category, _ := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
			parentCategory, _ := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: category.ParentID})
			// fmt.Println(i, parentCategory.Name, count)
			myMap[parentCategory.Name] += count
		}
	}
	for k, v := range myMap{
		fmt.Println(k, v)
	}
	
}

