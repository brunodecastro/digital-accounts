package input

// CreateAccountInputVO - vo that represents the input values of createAccount resource
type CreateAccountInputVO struct {
	CPF     string `json:"cpf" example:"008.012.461-56"`
	Name    string `json:"name" example:"Bruno de Castro Oliveira"`
	Secret  string `json:"secret" example:"123456"`
	Balance int    `json:"balance" example:"1000"`
}
