package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type orderRepo struct {
   fileName string
}

func NewOrderRepo(fileName string) *orderRepo {
	return &orderRepo{
		fileName:fileName,
	}
}



func (o *orderRepo) CreateOrder(order models.CreateOrder, total int) (string, error) {
	orderJson, err := ioutil.ReadFile(o.fileName)
	if err != nil {
		return "", err
	}
	
	ordersSlice := []models.Order{}

	json.Unmarshal(orderJson, &ordersSlice)

	uuid := uuid.New().String()

	newOrder := models.Order{
		OrderId: uuid,
		Cutomer_name: order.Cutomer_name,
		Customer_address: order.Customer_address,
		Customer_phone: order.Customer_phone,
		Total: total,
	}
	modifiedOrdersSlice := []models.Order{}
	modifiedOrdersSlice = append(modifiedOrdersSlice, newOrder)

	for _, v := range ordersSlice {
		modifiedOrdersSlice = append(modifiedOrdersSlice, v)
	}
	
	orderItemsJson, err := ioutil.ReadFile("./data/orderItems.json")
	if err != nil {
		return "", err
	}

	orderItems:=[]models.OrderItems{}
	json.Unmarshal(orderItemsJson, &orderItems)
	modifiedOrderItems := []models.OrderItems{}


	for _,v:=range order.OrderItems{
		orderItems=append(orderItems,models.OrderItems{
			Product_id: v.Product_id,
			Count: v.Count,
			Order_id: uuid,
		})
	}
	modifiedOrderItems = append(modifiedOrderItems, orderItems...)

	body, err := json.MarshalIndent(modifiedOrderItems, "", " ")
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile("./data/orderItems.json", body, os.ModePerm)
	if err != nil {
		return "", err
	}


	body, err = json.MarshalIndent(modifiedOrdersSlice, "", " ")
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(o.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}


func (o *orderRepo) GetByIdOrder(id string) (models.GetByIdOrder, error) {

	orderJson, err := ioutil.ReadFile(o.fileName)
	if err != nil {
		return models.GetByIdOrder{}, err
	}
	ordersSlice := []models.Order{}
	json.Unmarshal(orderJson, &ordersSlice)

	orderItemsJson, err := ioutil.ReadFile("./data/orderItems.json")
	if err != nil {
		return models.GetByIdOrder{}, err
	}
	orderItems:=[]models.OrderItems{}
	json.Unmarshal(orderItemsJson, &orderItems)
	response:=models.GetByIdOrder{}

	flag:=false
	for _, v := range ordersSlice {
		if id == v.OrderId {
			flag=true
			// response.Cutomer_name = v.Cutomer_name
			// response.Customer_address = v.Customer_address //----------
			// response.Customer_phone = v.Customer_address
		 			
		}
	}
	if !flag{
		return models.GetByIdOrder{}, errors.New("Order not found")
	}
	// for _, v := range orderItems{
	// 	if id == v.Order_id {
	// 		response.OrderItems=append(response.OrderItems,models.CreateOrderItems{
	// 			Product_id: v.Product_id,
	// 			Count: v.Count,
	// 		})	
	// 	}
	// }
	return response,nil
}