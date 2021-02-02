package output

// CreateAccountOutputVO - vo that represents the output values of createAccount resource
type CreateAccountOutputVO struct {
	CPF       string  `json:"cpf"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
