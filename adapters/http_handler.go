package adapters

import (
	"encoding/json"
	"net/http"

	"loans/core"
	"loans/ports"
)

// HTTPHandler é um handler HTTP que processa requisições relacionadas a empréstimos
type HTTPHandler struct {
	LoanService ports.LoanService
}

// CostumerLoans lida com a requisição do usuário
func (h *HTTPHandler) CostumerLoans(rw http.ResponseWriter, r* http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var client core.ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	loans := h.LoanService.GetLoans(client)

	response := map[string]interface{}{
		"costumer": client.Name,
		"loans": loans,
	}

	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)	
	}
}