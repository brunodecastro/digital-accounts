package vo

// CredentialVO - vo that represents the credentials of the user
type CredentialVO struct {
	CPF       string `json:"cpf"`
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
}
