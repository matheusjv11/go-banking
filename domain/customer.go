package domain

import (
	"github.com/matheusjv11/go-banking/dto"
	"github.com/matheusjv11/go-banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	Dateofbirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) statusAsText() string {
	if c.Status == "1" {
		return "active"
	}
	return "inactive"
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		Dateofbirth: c.Dateofbirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	// status == 1 status == = status = ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
