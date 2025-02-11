package core

// ClientRequest representa as informações que serão recebidas na requisição
type ClientRequest struct {
	Age 		uint64		`json:"age"`
	CPF 		string		`json:"cpf"`
	Name 		string		`json:"name"`
	Income 		float64		`json:"income"`
	Location 	string		`json:"location"`
}

// Loan representa as informações que serão retornadas como resposta
type Loan struct {
	Type 			string 	`json:"costumer"`
	InterestRate 	int 	`json:"interest_rate"`
}

// LoanService define um serviço para obtenção de empréstimos
type LoanService interface {
	GetLoans(client ClientRequest) []Loan
}

// LoanServiceImpl implementa interface LoanService
type LoanServiceImpl struct {}

// GetLoans implementa a lógica necessária para definir os empréstimos do usuário
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