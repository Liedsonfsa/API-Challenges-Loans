package ports

import "loans/core"

// LoanService interface utilizada pelo core para se comunicar com o mundo externo
type LoanService interface {
	GetLoans(client core.ClientRequest) []core.Loan
}