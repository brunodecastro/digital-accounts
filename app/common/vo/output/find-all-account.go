package output

// FindAllAccountOutputVO - vo that represents the output values of findAllAccounts
type FindAllAccountOutputVO struct {
	ID        string  `json:"id"`
	Cpf       string  `json:"cpf"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
