package input

// CreateTransferInputVO - vo that represents the input values of createTransfer resource
type CreateTransferInputVO struct {
	AccountOriginID      string
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int64  `json:"amount"`
}
