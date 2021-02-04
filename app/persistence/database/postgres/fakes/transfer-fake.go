package fakes

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/model"
	"time"
)

var (
	fakeTransfers = []model.Transfer{
		{
			ID:                   model.TransferID(common.NewUUID()),
			AccountOriginID:      GetFakeAccount1().ID,
			AccountDestinationID: GetFakeAccount2().ID,
			Amount:               100,
			CreatedAt:            time.Time{},
		},
		{
			ID:                   model.TransferID(common.NewUUID()),
			AccountOriginID:      GetFakeAccount1().ID,
			AccountDestinationID: GetFakeAccount1().ID,
			Amount:               200,
			CreatedAt:            time.Time{},
		},
	}
)

// GetFakeTransfers returns fake transfers
func GetFakeTransfers() *[]model.Transfer {
	return &fakeTransfers
}

// GenerateNewFakeTransfer generates a new fake transfer
func GenerateNewFakeTransfer() model.Transfer {
	return model.Transfer{
		ID:                   model.TransferID(common.NewUUID()),
		AccountOriginID:      GetFakeAccount1().ID,
		AccountDestinationID: GetFakeAccount2().ID,
		Amount:               100,
		CreatedAt:            time.Time{},
	}
}
