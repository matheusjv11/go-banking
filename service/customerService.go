package service

import (
	"github.com/matheusjv11/go-banking/domain"
	"github.com/matheusjv11/go-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repositoy domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositoy}
}
