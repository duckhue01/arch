package main

type OrderService struct {
	OrderRepo string
	Publisher DomainEventPublisher
}

func (cus *OrderService) TryCreateOrder() {

}

func (cus *OrderService) CommitCreateOrder() {

}

func (cus *OrderService) CancelCreateOrder() {

}