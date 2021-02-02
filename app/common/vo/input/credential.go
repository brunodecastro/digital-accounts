package input

// CredentialInputVO - vo that represents the input credential of the user
type CredentialInputVO struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}
