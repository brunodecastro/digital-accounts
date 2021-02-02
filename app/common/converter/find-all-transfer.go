package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// ModelToFindAllTransferOutputVO - converts []model.Transfer to []output.FindAllTransferOutputVO
func ModelToFindAllTransferOutputVO(transfers []model.Transfer) []output.FindAllTransferOutputVO {
	var transfersOutputVO = make([]output.FindAllTransferOutputVO, 0)
	for _, transfer := range transfers {
		transfersOutputVO = append(transfersOutputVO, output.FindAllTransferOutputVO{
			ID:                   string(transfer.ID),
			AccountOriginID:      string(transfer.AccountOriginID),
			AccountDestinationID: string(transfer.AccountDestinationID),
			Amount:               transfer.Amount.ToFloat64(),
			CreatedAt:            util.FormatDate(transfer.CreatedAt),
		})
	}
	return transfersOutputVO
}
