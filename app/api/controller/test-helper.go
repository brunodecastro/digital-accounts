package controller

import (
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"time"
)

// Mocks a handler and returns a httptest.ResponseRecorder
func mockRequestHandler(
	req *http.Request,
	method string,
	strPath string,
	authRequest bool,
	fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rec := httptest.NewRecorder()

	if authRequest {
		// Creates a token by credential
		token := auth.GenerateToken(vo.CredentialClaimsVO{
			Username:  "Bruno",
			AccountId: "0001",
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		})
		req.Header.Set("Authorization", "Bearer "+token)
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rec, req)
	return rec
}

// Mocks a handler func and returns a httptest.ResponseRecorder
func mockRequestHandlerFunc(req *http.Request, method string, strPath string, authRequest bool, handlerFunc http.HandlerFunc) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.HandlerFunc(method, strPath, handlerFunc)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rec := httptest.NewRecorder()

	if authRequest {
		// Creates a token by credential
		token := auth.GenerateToken(vo.CredentialClaimsVO{
			Username:  "Bruno",
			AccountId: "0001",
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		})
		req.Header.Set("Authorization", "Bearer "+token)
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rec, req)
	return rec
}
