package core

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

type LoanService interface {
	GetLoans(client ClientRequest) []Loan
}

type LoanServiceImpl struct {}

func (s *LoanServiceImpl) GetLoans(client ClientRequest) []Loan {
	var loans []Loan

	if client.Income <= 3000{
		loans = append(loans, Loan{Type: "PERSONAL", InterestRate: 4})
		loans = append(loans, Loan{Type: "GUARANTEED", InterestRate: 3})
	} else if (client.Income > 3000 && client.Income <= 5000) && client.Age < 30 && client.Location == "SP" {
		loans = append(loans, Loan{Type: "PERSONAL", InterestRate: 4})
		loans = append(loans, Loan{Type: "GUARANTEED", InterestRate: 3})
	}
	
	if client.Income >= 5000 {
		loans = append(loans, Loan{Type: "CONSIGNMENT", InterestRate: 2})
	}

	return loans
}