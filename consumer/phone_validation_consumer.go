package consumer

import (
	"fmt"

	"github.com/dayvsonspacca/go-phone-validation/handler"
)

func InitializePhoneValidationConsumer(phoneValidationQueue chan handler.PhoneValidationData) {
	fmt.Println("PhoneValidationConsume started.")

	go func() {
		for {
			handler.HandlerPhoneValidation(<-phoneValidationQueue)
			fmt.Printf("Total in queue: %d\n", len(phoneValidationQueue))
		}
	}()

	select {}
}
