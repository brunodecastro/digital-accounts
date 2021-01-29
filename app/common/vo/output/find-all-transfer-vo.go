package output

type FindAllTransferOutputVO struct {
	Id                   string  `json:"id"`
	AccountOriginID      string  `json:"account_origin_id"`
	AccountDestinationID string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
	CreatedAt            string  `json:"created_at"`
}
