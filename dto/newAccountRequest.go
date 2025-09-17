package dto

import (
	"strings"

	"github.com/matheusjv11/go-banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("amount must be at least 5000")
	}

	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("account type must be savings or checking")
	}

	return nil
}
