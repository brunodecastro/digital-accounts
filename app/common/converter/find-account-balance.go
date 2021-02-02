package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// ModelToFindAccountBalanceOutputVO - converts model.Account to input.FindAccountBalanceOutputVO
func ModelToFindAccountBalanceOutputVO(account *model.Account) output.FindAccountBalanceOutputVO {
	return output.FindAccountBalanceOutputVO{
		ID:      string(account.ID),
		Balance: account.Balance.ToFloat64(),
	}
}
