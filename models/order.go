package models


type Order struct{
	OrderId string	`json:"orederId"`
	Cutomer_name string `json:"cutomer_name"`
	Customer_address string `json:"customer_address"`
	Customer_phone string `json:"customer_phone"`
	Total int `json:"total"`
  }
  
  type CreateOrder struct{
	Cutomer_name string `json:"cutomer_name"`
	Customer_address string `json:"customer_address"`
	Customer_phone string `json:"customer_phone"`
	OrderItems []CreateOrderItems `json:"order_items"`
	
  }
  type CreateOrderItems struct{
	Product_id string `json:"product_id"`
	Count int `json:"count"`		
  }
  
  type OrderItems struct{
	Product_id string `json:"product_id"`
	Count int `json:"count"`
	Order_id string `json:"order_id"`
  }

  type GetByIdOrder struct{
	OrderId string	`json:"orederId"`
	Cutomer_name string `json:"cutomer_name"`
	Customer_address string `json:"customer_address"`
	Customer_phone string `json:"customer_phone"`
	Total int `json:"total"`
	OrderItems []CreateOrderItems `json:"order_items"`
  }