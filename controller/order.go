package controller

import (
	"app/models"
)


func (c *Controller) CreateOrder(order models.CreateOrder) (string, error){
	
	var sum int
	for _, v := range order.OrderItems{
		
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.Product_id })
		if err != nil {
			return "", nil
		}
		sum += int(product.Price) * v.Count
	}
	
	id, err := c.store.Order().CreateOrder(order, sum)
	if err != nil {
		return "", err
	}	
	return id, nil
}

func (c *Controller) GetByIdOrder(id string) (models.GetByIdOrder, error){
	res, err := c.store.Order().GetByIdOrder(id)
	if err != nil {
		return models.GetByIdOrder{}, err
	}	
	return res, nil
}