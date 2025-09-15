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
	// status == 1 status == = status = ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
