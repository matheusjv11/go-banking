package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusjv11/go-banking/dto"
	"github.com/matheusjv11/go-banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
