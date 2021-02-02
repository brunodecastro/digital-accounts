package input

// CreateAccountInputVO - vo that represents the input values of createAccount resource
type CreateAccountInputVO struct {
	CPF     string `json:"cpf"`
	Name    string `json:"name"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}
