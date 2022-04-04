package main

type CustomerService struct {
	OrderRepo string
	Publisher DomainEventPublisher
}

func (cus *CustomerService) TryCreateOrder() {

}

func (cus *CustomerService) CommitCreateOrder() {

}

func (cus *CustomerService) CancelCreateOrder() {

}