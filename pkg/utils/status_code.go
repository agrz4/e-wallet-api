package utils

import (
	"e-wallet-api-go/pkg/customerror"
	"net/http"
)

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*customerror.NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*customerror.UserAlreadyExistError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*customerror.IncorrectCredentialsError); ok {
		statusCode = http.StatusUnauthorized
	} else if _, ok := err.(*customerror.UserNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*customerror.PasswordNotSame); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*customerror.ResetTokenNotFound); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*customerror.SourceOfFundNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*customerror.InsufficientBalanceError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*customerror.WalletNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*customerror.WalletAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*customerror.TransferToSameWalletError); ok {
		statusCode = http.StatusBadRequest
	}
	return statusCode
}
