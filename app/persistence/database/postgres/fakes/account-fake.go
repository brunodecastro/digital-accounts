package fakes

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

var (
	fakeAccounts = []model.Account{
		{
			ID:        model.AccountID(common.NewUUID()),
			CPF:       util.GenerateCPFOnlyNumbers(),
			Name:      "Fake 1",
			Secret:    util.EncryptPassword("secret"),
			Balance:   types.Money(10000),
			CreatedAt: time.Time{},
		},
		{
			ID:        model.AccountID(common.NewUUID()),
			CPF:       util.GenerateCPFOnlyNumbers(),
			Name:      "Fake 2",
			Secret:    util.EncryptPassword("secret"),
			Balance:   types.Money(10000),
			CreatedAt: time.Time{},
		},
	}
)

// GetFakeAccounts returns the accounts fake
func GetFakeAccounts() *[]model.Account {
	return &fakeAccounts
}

// GetFakeAccount1 returns fake account represents by 1
func GetFakeAccount1() model.Account {
	return fakeAccounts[0]
}

// GetFakeAccount2 returns fake account represents by 1
func GetFakeAccount2() model.Account {
	return fakeAccounts[1]
}

// GenerateNewFakeAccount generates a new account fake
func GenerateNewFakeAccount(name string) model.Account {
	return model.Account{
		ID:        model.AccountID(common.NewUUID()),
		CPF:       util.GenerateCPFOnlyNumbers(),
		Name:      name,
		Secret:    util.EncryptPassword("secret"),
		Balance:   types.Money(10000),
		CreatedAt: time.Time{},
	}
}
