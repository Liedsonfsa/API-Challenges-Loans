package ports

import "loans/core"

type LoanService interface {
	GetLoans(client core.ClientRequest) []core.Loan
}