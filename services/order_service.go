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
	orderRepository repositories.IOrder
}

func NewOrderService(repository repositories.IOrder) IOrderService {
	return &OrderService{repository}
}

func (o *OrderService) GetOrderByID(orderID int64) (*datamodels.Order, error) {
	return o.orderRepository.SelectByKey(orderID)
}

func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.orderRepository.SelectAll()
}

func (o *OrderService) DeleteOrderByID(orderID int64) bool {
	return o.orderRepository.Delete(orderID)
}

func (o *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return o.orderRepository.Insert(order)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.orderRepository.Update(order)
}

func (o *OrderService) GetAllOrderWithInfo() (map[int]map[string]string, error) {
	return o.orderRepository.SelectAllWithInfo()
}
