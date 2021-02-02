package converter

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

// CreateTransferInputVOToModel - converts input.CreateTransferInputVO to model.Transfer
func CreateTransferInputVOToModel(transferInputVO input.CreateTransferInputVO) model.Transfer {
	return model.Transfer{
		ID:                   model.TransferID(common.NewUUID()),
		AccountOriginID:      model.AccountID(transferInputVO.AccountOriginID),
		AccountDestinationID: model.AccountID(transferInputVO.AccountDestinationID),
		Amount:               types.Money(transferInputVO.Amount),
		CreatedAt:            time.Now(),
	}
}

// ModelToCreateTransferOutputVO - converts model.Transfer to input.CreateTransferOutputVO
func ModelToCreateTransferOutputVO(transfer *model.Transfer) output.CreateTransferOutputVO {
	return output.CreateTransferOutputVO{
		ID:                   string(transfer.ID),
		AccountOriginID:      string(transfer.AccountOriginID),
		AccountDestinationID: string(transfer.AccountDestinationID),
		Amount:               types.Money(transfer.Amount).ToFloat64(),
		CreatedAt:            util.FormatDate(transfer.CreatedAt),
	}
}
