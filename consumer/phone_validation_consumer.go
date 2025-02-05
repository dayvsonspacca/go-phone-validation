package consumer

import (
	"github.com/dayvsonspacca/go-phone-validation/handler"
	"github.com/dayvsonspacca/go-phone-validation/model"
)

func InitializePhoneValidationConsumer(phoneValidationQueue chan model.PhoneValidationRequest) {
	go func() {
		for {
			phoneValidationRequest := <-phoneValidationQueue
			result := handler.ProcessPhoneValidationRequest(phoneValidationRequest.Phone, phoneValidationRequest.NationalIdentyNumber)

			handler.SendPhoneValidationResponse(phoneValidationRequest.CallbackUrl, handler.PhoneValidationResponse{
				Token:  phoneValidationRequest.Token,
				Result: result,
			})
		}
	}()

	select {}
}
