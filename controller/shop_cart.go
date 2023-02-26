package controller

import (
	"app/models"
	"errors"
	"sort"
	"time"
)

func (c *Controller) AddShopCart(req *models.Add) (string, error) {
	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return "", err
	}
	
	id, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return err
}

func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
	_, err := c.store.User().GetByID(req)
	if err != nil {
		return 0, err
	}
	
	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, v := range users {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		if status == "fixed" {
			total += float64(v.Count) * (product.Price - discount)
		} else if status == "percent" {
			if discount < 0 || discount > 100 {
				return 0, errors.New("Invalid discount range")
			}
			total += float64(v.Count) * (product.Price - (product.Price * discount)/100)
		} else {
			return 0, errors.New("Invalid status name")
		}
	}

	if total < 0 {
		return 0, nil
	}
	return total, nil
}


func (c *Controller) ShopCartDataFilter (fromDate, toDate string) ([]models.ShopCart, error) {
	shopCarts, err := c.store.ShopCart().Read()
	if err != nil {
		return []models.ShopCart{}, err
	}

	defaultTimeLayout := "2006-01-02 15:04:05"
	dateSlice := []time.Time{}

	if fromDate == "" && toDate == "" {
		for _, v := range shopCarts{
			dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
			  dateSlice = append(dateSlice, dataInShopCart)
		  }
	}else{
	modFromDate := fromDate + " 00:00:00"
	modToDate := toDate + " 23:59:59"

	strart, err1 := time.Parse(defaultTimeLayout, modFromDate)
	if err1 != nil {
		return []models.ShopCart{}, err1
	}
	end, err2 := time.Parse(defaultTimeLayout, modToDate)
	if err2 != nil {
		return []models.ShopCart{}, err2
	}
	for _, v := range shopCarts{
	  dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
	  if dataInShopCart.After(strart) && dataInShopCart.Before(end) {
		dateSlice = append(dateSlice, dataInShopCart)
	  }
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


	sortedSlice := []models.ShopCart{}	
	for _, sortedData := range dateSlice{
		for _, v := range shopCarts{
			dataInShopCart, _ := time.Parse(defaultTimeLayout, v.Time)
			if sortedData == dataInShopCart {
				// fmt.Println(v)
				sortedSlice = append(sortedSlice, v)
			}
	}
	}
	
	return sortedSlice, err
}