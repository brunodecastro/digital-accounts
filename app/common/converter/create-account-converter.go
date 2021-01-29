package converter

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

func CreateAccountInputVOToModel(accountInputVO input.CreateAccountInputVO) model.Account {
	return model.Account{
		Id:        model.AccountID(common.NewUUID()),
		Name:      accountInputVO.Name,
		Cpf:       util.NumbersOnly(accountInputVO.Cpf),
		Secret:    util.EncryptPassword(accountInputVO.Secret),
		Balance:   common.Money(accountInputVO.Balance),
		CreatedAt: time.Now(),
	}
}

func ModelToCreateAccountOutputVO(account *model.Account) output.CreateAccountOutputVO {
	return output.CreateAccountOutputVO{
		Name:      account.Name,
		Cpf:       util.FormatCpf(account.Cpf),
		Balance:   account.Balance.GetFloat(),
		CreatedAt: util.FormatDate(account.CreatedAt),
	}
}
