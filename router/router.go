package router

import (
	"net/http"

	"github.com/dayvsonspacca/go-phone-validation/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Initialize(phoneValidationQueue chan model.PhoneValidationRequest) *gin.Engine {
	router := gin.New()

	router.POST("/api/v1/validate-phone-number", func(ctx *gin.Context) {
		var phoneValidationRequest model.PhoneValidationRequest

		if err := ctx.ShouldBindJSON(&phoneValidationRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		phoneValidationRequest.Token = uuid.New().String()
		// [ TODO ]: Validate fields

		phoneValidationQueue <- phoneValidationRequest

		ctx.JSON(200, gin.H{
			"token": phoneValidationRequest.Token,
		})
	})

	return router
}
