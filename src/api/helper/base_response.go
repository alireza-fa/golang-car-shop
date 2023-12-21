package helper

import "github.com/alireza-fa/golang-car-shop/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       ResultCode                     `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode ResultCode, err any) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}
