package output

type CredentialOutputVO struct {
	Cpf       string `json:"cpf"`
	AccountId string `json:"account_id"`
	Username  string `json:"username"`
	Token     string `json:"token"`
}
