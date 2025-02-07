package router

import (
	"fmt"
	"net/http"

	"github.com/dayvsonspacca/go-phone-validation/handler"
	"github.com/dayvsonspacca/go-phone-validation/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Initialize(phoneValidationQueue chan handler.PhoneValidationData) *gin.Engine {
	router := gin.New()

	router.POST("/api/v1/validate-phone-number", func(ctx *gin.Context) {
		phoneValidationRequest, err := request.ParsePhoneValidationRequest(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		token := uuid.New().String()

		phoneValidationQueue <- handler.PhoneValidationData{
			PhoneValidationRequest: phoneValidationRequest,
			Token:                  token,
		}

		fmt.Printf("New PhoneValidationData added to queue, token: %s\n", token)

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	router.POST("/api/v1/new-validated-phone-number", func(ctx *gin.Context) {
		newPhoneValidatedRequest, err := request.ParseNewPhoneValidatedRequest(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		id := handler.HandlerNewValidatedPhoneNumber(newPhoneValidatedRequest)
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})
	return router
}
