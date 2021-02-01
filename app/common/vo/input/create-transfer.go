package input

type CreateTransferInputVO struct {
	AccountOriginId      string
	AccountDestinationId string `json:"account_destination_id"`
	Amount               int64  `json:"amount"`
}
