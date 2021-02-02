package output

// FindAccountBalanceOutputVO - vo that represents the output values of findAccountBalance
type FindAccountBalanceOutputVO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}
