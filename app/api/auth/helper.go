package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"net/http"
)

// GetAccountIdFromAuth return the account id from the auth token
func GetAccountIdFromToken(req *http.Request) string {
	// Get the credential claims from the context
	credentialClaimsVO := req.Context().Value(constants.CredentialClaimsContextKey).(vo.CredentialClaimsVO)
	return credentialClaimsVO.AccountId
}
