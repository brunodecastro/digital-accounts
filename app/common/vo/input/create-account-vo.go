package input

type CreateAccountInputVO struct {
	Cpf     string `json:"cpf" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Secret  string `json:"secret" validate:"required"`
	Balance int    `json:"balance"`
}
