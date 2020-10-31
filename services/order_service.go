package services

import (
	"spike/datamodels"
	"spike/repositories"
)

type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderWithInfo() (map[int]map[string]string, error)
	DeleteOrderByID(int64) bool
	InsertOrder(*datamodels.Order) (int64, error)
	UpdateOrder(*datamodels.Order) error
}

type OrderService struct {
	OrderRepository repositories.IOrder
}

func NewOrderService(repository repositories.IOrder) IOrderService {
	return &OrderService{repository}
}

func (o *OrderService) GetOrderByID(orderID int64) (order *datamodels.Order, err error) {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) DeleteOrderByID(orderID int64) (isOk bool) {
	return o.OrderRepository.Delete(orderID)
}

func (o *OrderService) InsertOrder(order *datamodels.Order) (orderID int64, err error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) GetAllOrderWithInfo() (map[int]map[string]string, error) {
	return o.OrderRepository.SelectAllWithInfo()
}
