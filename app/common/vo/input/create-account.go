package input

type CreateAccountInputVO struct {
	Cpf     string `json:"cpf"`
	Name    string `json:"name"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}
