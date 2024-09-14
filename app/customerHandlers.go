package app

import (
	"HexagonalProject/domain/errs"
	"HexagonalProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	status := queryParams.Get("status")
	customers, err := ch.getCustomers(status)
	ch.CustomerService.GetAllCustomer()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandler) getCustomers(status string) (interface{}, *errs.AppError) {
	if status == "" {
		return ch.CustomerService.GetAllCustomer()
	} else if (status == "active") || (status == "inactive") {
		return ch.CustomerService.GetAllActiveorInactiveCustomers(status == "active")
	} else {
		return nil, errs.NewNotFoundError("Unexpected Status Found")
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := ch.CustomerService.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
