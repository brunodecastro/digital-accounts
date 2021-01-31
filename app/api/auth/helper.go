package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/gorilla/context"
	"net/http"
)

// GetAccountIdFromAuth return the account id from the auth token
func GetAccountIdFromToken(req *http.Request) string {
	credentialClaimsVO := context.Get(req, "credentialClaims").(vo.CredentialClaimsVO)
	return credentialClaimsVO.AccountId
}
