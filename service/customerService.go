package service

import "github.com/matheusjv11/go-banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repositoy domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositoy}
}
