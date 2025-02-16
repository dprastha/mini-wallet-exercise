package utils

import (
	"mini-wallet-exercise/pkg"
)

func SuccessResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		Status: "success",
		Data:   &data,
	})
}

func ErrorResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		Status: "fail",
		Data:   &data,
	})
}

func ErrorValidationResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		Status: "fail",
		Data:   &data,
	})
}
