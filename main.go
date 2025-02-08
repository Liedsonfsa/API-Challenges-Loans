package main

import (
	"loans/adapters"
	"loans/core"
	"net/http"
)

func main() {

	loanService := &core.LoanServiceImpl{}
	handler := &adapters.HTTPHandler{LoanService: loanService}

	http.HandleFunc("/costumer-loans", handler.CostumerLoans)

	http.ListenAndServe(":3000", nil)
}