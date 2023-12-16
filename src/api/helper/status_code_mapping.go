package helper

import (
	"github.com/alireza-fa/golang-car-shop/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// OTP
	service_errors.OtpExists:   409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
