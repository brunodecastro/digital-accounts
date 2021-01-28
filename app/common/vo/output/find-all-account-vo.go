package input

type FindAllAccountOutputVO struct {
	Id        string  `json:"id"`
	Cpf       string  `json:"cpf"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
