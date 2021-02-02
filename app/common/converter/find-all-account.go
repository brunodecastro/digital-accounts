package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// AccountModelToFindAllAccountOutputVO - converts []model.Account to []output.FindAllAccountOutputVO
func AccountModelToFindAllAccountOutputVO(accounts []model.Account) []output.FindAllAccountOutputVO {
	var accountsOutputVO = make([]output.FindAllAccountOutputVO, 0)
	for _, account := range accounts {
		accountsOutputVO = append(accountsOutputVO, output.FindAllAccountOutputVO{
			ID:        string(account.ID),
			Cpf:       util.FormatCpf(account.CPF),
			Name:      account.Name,
			Balance:   account.Balance.ToFloat64(),
			CreatedAt: util.FormatDate(account.CreatedAt),
		})
	}
	return accountsOutputVO
}
