package converter

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

func CreateTransferInputVOToModel(transferInputVO input.CreateTransferInputVO) model.Transfer {
	return model.Transfer{
		Id:                   model.TransferID(common.NewUUID()),
		AccountOriginId:      model.AccountID(transferInputVO.AccountOriginId),
		AccountDestinationId: model.AccountID(transferInputVO.AccountDestinationId),
		Amount:               common.Money(transferInputVO.Amount),
		CreatedAt:            time.Now(),
	}
}

func ModelToCreateTransferOutputVO(transfer *model.Transfer) output.CreateTransferOutputVO {
	return output.CreateTransferOutputVO{
		Id:                   string(transfer.Id),
		AccountOriginID:      string(transfer.AccountOriginId),
		AccountDestinationID: string(transfer.AccountDestinationId),
		Amount:               common.Money(transfer.Amount).GetFloat(),
		CreatedAt:            util.FormatDate(transfer.CreatedAt),
	}
}
