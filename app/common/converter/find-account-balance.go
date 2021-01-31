package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
)

func ModelToFindAccountBalanceOutputVO(account *model.Account) output.FindAccountBalanceOutputVO {
	return output.FindAccountBalanceOutputVO{
		Id:      string(account.Id),
		Balance: account.Balance.GetFloat64(),
	}
}
