package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type NewPhoneValidatedRequest struct {
	Phone                string `json:"phone" binding:"required"`
	NationalIdentyNumber string `json:"national_identy_number" binding:"required"`
	Status               string `json:"status" binding:"required"`
}

func ParseNewPhoneValidatedRequest(ctx *gin.Context) (NewPhoneValidatedRequest, error) {
	var newPhoneValidatedRequest NewPhoneValidatedRequest

	if err := ctx.ShouldBindJSON(&newPhoneValidatedRequest); err != nil {
		// [ TODO ] Enhance error messages
		return NewPhoneValidatedRequest{}, fmt.Errorf("error in NewPhoneValidatedRequest: %v", err)
	}

	// [ TODO ] Add fields validation

	return newPhoneValidatedRequest, nil
}
