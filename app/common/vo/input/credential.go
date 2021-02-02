package input

// CredentialInputVO - vo that represents the input credential of the user
type CredentialInputVO struct {
	CPF    string `json:"cpf" example:"00801246156"`
	Secret string `json:"secret" example:"123456"`
}
