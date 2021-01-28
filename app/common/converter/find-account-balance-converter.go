package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
)

func AccountModelToFindAccountBalanceOutputVO(account *model.Account) output.FindAccountBalanceOutputVO {
	return output.FindAccountBalanceOutputVO{
		Id:      string(account.ID),
		Balance: account.Balance.GetFloat(),
	}
}
