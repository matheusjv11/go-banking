package service

import (
	"github.com/matheusjv11/go-banking/domain"
	"github.com/matheusjv11/go-banking/dto"
	"github.com/matheusjv11/go-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repositoy domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositoy}
}
