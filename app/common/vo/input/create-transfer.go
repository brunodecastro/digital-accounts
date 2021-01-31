package input

type CreateTransferInputVO struct {
	AccountOriginId      string
	AccountDestinationId string `json:"account_destination_id" validate:"required"`
	Amount               int64  `json:"amount" validate:"gt=0,required"`
}
