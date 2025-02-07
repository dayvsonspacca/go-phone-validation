package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PhoneValidationRequest struct {
	Phone                string `json:"phone" binding:"required"`
	NationalIdentyNumber string `json:"national_identy_number" binding:"required"`
	CallbackUrl          string `json:"callback_url" binding:"required"`
}

func ParsePhoneValidationRequest(ctx *gin.Context) (PhoneValidationRequest, error) {
	var phoneValidationRequest PhoneValidationRequest

	if err := ctx.ShouldBindJSON(&phoneValidationRequest); err != nil {
		// [ TODO ] Enhance error messages
		return PhoneValidationRequest{}, fmt.Errorf("error in ParsePhoneValidationRequest: %v", err)
	}

	// [ TODO ] Add fields validation

	return phoneValidationRequest, nil
}
