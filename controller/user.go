package controller

import (
	"app/models"
	"errors"
	"fmt"
)

func (c *Controller) CreateUser(req *models.CreateUser) (string, error) {
	id, err := c.store.User().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteUser(req *models.UserPrimaryKey) error {
	err := c.store.User().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateUser(req *models.UpdateUser, userId string) error {
	err := c.store.User().Update(req, userId)
	if err != nil {
		return err
	}
	return nil
} 

func (c *Controller) GetByIdUser(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetByID(req)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (c *Controller) GetAllUser(req *models.GetListRequest) (models.GetListResponse, error) {
	users, err := c.store.User().GetAll(req)
	if err != nil {
		return models.GetListResponse{}, err
	}
	return users, nil
}

func (c *Controller) WithdrawCheque(total float64, userId string) error {
	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})
	if err != nil {
		return err
	}
	if user.Balance >= total {
		user.Balance -= total
	} else {
		return errors.New("You don't have enough money")
	}

	err = c.store.User().Update(&models.UpdateUser{
		Balance: user.Balance,
	}, userId)
	if err != nil {
		return err
	}

	err = c.store.ShopCart().UpdateShopCart(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) MoneyTransfer(sender string, receiver string, money float64) error {
	send, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: sender}) 
	if err != nil {
		return err
	}

	receive, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: receiver}) 
	if err != nil {
		return err
	}

	comMoney := 0.1 * float64(money)
	if money+comMoney > send.Balance {
		return errors.New("Sender doesn't have enough money")
	}
	send.Balance -= money + comMoney
	err = c.store.User().Update(&models.UpdateUser{
		Name: send.Name,
		Surname: send.Surname,
		Balance: send.Balance,
	}, sender)
	if err != nil {
		return err
	}

	err = c.store.Commission().AddCommission(&models.Commission{
		Balance: comMoney,
	})

	receive.Balance += money
	err = c.store.User().Update(&models.UpdateUser{
		Name: receive.Name,
		Surname: receive.Surname,
		Balance: receive.Balance,
	}, receiver)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) ClientHistory (req string){
	clients, err :=c.store.User().Read()
	if err != nil {
		fmt.Println(err)
	}	
	
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
	
	for _, clientsV := range clients{
		if req == clientsV.Name {
			fmt.Println("History of client: ",clientsV.Name, clientsV.Surname)
			i := 0
			for _, shopCartsV := range shopCarts{
				if clientsV.Id == shopCartsV.UserId && shopCartsV.Status {
					i++
					products, _ := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopCartsV.ProductId})
					fmt.Println(i, "Product name:",products.Name, "  Price:", products.Price, "  Count:", shopCartsV.Count,
					"  Total:", int(products.Price)*shopCartsV.Count, "  Time:", shopCartsV.Time)
				}
			}
		}
	}
}


func (c *Controller) ClientTotalMoneySpend(req string){
	clients, err :=c.store.User().Read()
	if err != nil {
		fmt.Println(err)
	}	
	
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
		
	for _, clientsV := range clients{
		if req == clientsV.Name {
			sum := 0
			for _, shopCartsV := range shopCarts{
				if clientsV.Id == shopCartsV.UserId && shopCartsV.Status {
					products, _ := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopCartsV.ProductId})
					sum += int(products.Price)*shopCartsV.Count
				}
			}
			fmt.Println(clientsV.Name, clientsV.Surname, " Total Buy Price:", sum)
		}
	}
}


func (c *Controller) MostActiveClient(){
	clients, err :=c.store.User().Read()
	if err != nil {
		fmt.Println(err)
	}	
	
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
	myMap := make(map[int]string)	
	for _, clientsV := range clients{
			sum := 0
			for _, shopCartsV := range shopCarts{
				if clientsV.Id == shopCartsV.UserId && shopCartsV.Status {
					products, _ := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopCartsV.ProductId})
					sum += int(products.Price)*shopCartsV.Count
				}
			}
			myMap[sum] = clientsV.Name + " " + clientsV.Surname
	}
	lk := 0 
	for k := range myMap{
		if lk < k {
			lk, k = k, lk
		}	
	}
	fmt.Println("The most active client is:", myMap[lk], "with spend money", lk)
}

func (c *Controller) DiscountToClient(req string){
	clients, err :=c.store.User().Read()
	if err != nil {
		fmt.Println(err)
	}	
	
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		fmt.Println(err)
	}
	totalCount, totalPrice, discount := 0, 0, 0 
	for _, clientsV := range clients{
		if req == clientsV.Name {
			fmt.Println("Discount to client: ",clientsV.Name, clientsV.Surname)
			for _, shopCartsV := range shopCarts{
				if clientsV.Id == shopCartsV.UserId && shopCartsV.Status {
					products, _ := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopCartsV.ProductId})
					
					totalCount += shopCartsV.Count
					totalPrice += int(products.Price)*shopCartsV.Count
					discount = int(products.Price)

					fmt.Println("Product name:",products.Name, "  Price:", products.Price, "  Count:", shopCartsV.Count,
					"  Total:", int(products.Price)*shopCartsV.Count, "  Time:", shopCartsV.Time)
				}
			}
			if totalCount > 9 {
			fmt.Println("Total count of products:", totalCount, "Total price of products:", totalPrice)
			fmt.Println("Your discount is:", discount, "total price with discount:", totalPrice-discount)	
			}else{
			fmt.Println("Total count of products:", totalCount, "Total price of products:", totalPrice)
			fmt.Println("Please buy more than 9 products to have discount")
			}
		}
	}
}