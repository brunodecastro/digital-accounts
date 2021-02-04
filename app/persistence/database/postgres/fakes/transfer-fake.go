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
			AccountOriginID:      "",
			AccountDestinationID: "",
			Amount:               0,
			CreatedAt:            time.Time{},
		},
		{
			ID:                   model.TransferID(common.NewUUID()),
			AccountOriginID:      "",
			AccountDestinationID: "",
			Amount:               0,
			CreatedAt:            time.Time{},
		},
	}
)