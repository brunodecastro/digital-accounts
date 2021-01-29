package input

type 	CreateTransferInputVO struct {
	AccountOriginId      string `json:"account_origin_id" validate:"required"`
	AccountDestinationId string `json:"account_destination_id" validate:"required"`
	Amount               int    `json:"amount" validate:"gt=0,required"`
}
