package converter

import (
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
)

func AccountModelToFindAllAccountOutputVO(accounts []model.Account) []output.FindAllAccountOutputVO {
	var accountsOutputVO = make([]output.FindAllAccountOutputVO, 0)
	for _, account := range accounts {
		accountsOutputVO = append(accountsOutputVO, output.FindAllAccountOutputVO{
			Id:        string(account.ID),
			Cpf:       util.FormatCpf(account.Cpf),
			Name:      account.Name,
			Balance:   account.Balance.GetFloat(),
			CreatedAt: util.FormatDate(account.CreatedAt),
		})
	}
	return accountsOutputVO
}
