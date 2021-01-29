package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
)

func ModelToFindAllTransferOutputVO(transfers []model.Transfer) []output.FindAllTransferOutputVO {
	var transfersOutputVO = make([]output.FindAllTransferOutputVO, 0)
	for _, transfer := range transfers {
		transfersOutputVO = append(transfersOutputVO, output.FindAllTransferOutputVO{
			Id:                   string(transfer.Id),
			AccountOriginID:      string(transfer.AccountOriginId),
			AccountDestinationID: string(transfer.AccountDestinationId),
			Amount:               transfer.Amount.GetFloat(),
			CreatedAt:            util.FormatDate(transfer.CreatedAt),
		})
	}
	return transfersOutputVO
}
