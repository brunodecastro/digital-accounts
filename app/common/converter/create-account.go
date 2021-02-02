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

// CreateAccountInputVOToModel - converts input.CreateAccountInputVO to model.Account
func CreateAccountInputVOToModel(accountInputVO input.CreateAccountInputVO) model.Account {
	return model.Account{
		ID:        model.AccountID(common.NewUUID()),
		Name:      accountInputVO.Name,
		Cpf:       util.NumbersOnly(accountInputVO.Cpf),
		Secret:    util.EncryptPassword(accountInputVO.Secret),
		Balance:   types.Money(accountInputVO.Balance),
		CreatedAt: time.Now(),
	}
}

// ModelToCreateAccountOutputVO - converts model.Account to input.CreateAccountInputVO
func ModelToCreateAccountOutputVO(account *model.Account) output.CreateAccountOutputVO {
	return output.CreateAccountOutputVO{
		Name:      account.Name,
		Cpf:       util.FormatCpf(account.Cpf),
		Balance:   account.Balance.ToFloat64(),
		CreatedAt: util.FormatDate(account.CreatedAt),
	}
}
