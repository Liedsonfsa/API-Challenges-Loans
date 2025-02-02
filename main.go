package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ClientRequest struct {
	Age 		uint64		`json:"age"`
	CPF 		string		`json:"cpf"`
	Name 		string		`json:"name"`
	Income 		float64		`json:"income"`
	Location 	string		`json:"location"`
}

type Loan struct {
	Type 			string 	`json:"costumer"`
	InterestRate 	int 	`json:"interest_rate"`
}

type Response struct {
	Costumer 	string `json:"costumer"`
	Loans 		[]Loan `json:"loans"`
}

func costumerLoans(w http.ResponseWriter, r* http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var client ClientRequest
	err = json.Unmarshal(body, &client)
	if err != nil {
		panic(err)
	}

	loans := getCostumer(client)

	response := Response{
		Costumer: 	client.Name,
		Loans: 		loans,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getCostumer(client ClientRequest) []Loan {
	var loans []Loan

	if client.Income <= 3000{
		loans = append(loans, Loan{Type: "PERSONAL", InterestRate: 4})
		loans = append(loans, Loan{Type: "GUARANTEED", InterestRate: 3})
	} else if (client.Income > 3000 && client.Income <= 5000) && client.Age < 30 && strings.EqualFold(client.Location, "SP") {
		loans = append(loans, Loan{Type: "PERSONAL", InterestRate: 4})
		loans = append(loans, Loan{Type: "GUARANTEED", InterestRate: 3})
	}

	if client.Income >= 5000 {
		loans = append(loans, Loan{Type: "CONSIGNMENT", InterestRate: 2})
	}

	return loans
}

func main() {

	http.HandleFunc("/costumer-loans", func(w http.ResponseWriter, r* http.Request) {
		switch r.URL.Path {
		case "/costumer-loans":
			switch r.Method {
			case "GET":
				w.Write([]byte("Hello World!"))
			case "POST":
				costumerLoans(w, r)
			default:
				w.Write([]byte("Esse método não existe"))
			}
		default:
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":3000", nil)
}