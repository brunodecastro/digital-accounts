package input

type FindAllAccountOutputVO struct {
	Cpf       string  `json:"cpf""`
	Name      string  `json:"name""`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
