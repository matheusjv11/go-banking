package domain

import "github.com/matheusjv11/go-banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	Dateofbirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
