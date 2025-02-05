package main

import (
	"github.com/dayvsonspacca/go-phone-validation/consumer"
	"github.com/dayvsonspacca/go-phone-validation/model"
	"github.com/dayvsonspacca/go-phone-validation/router"
)

func main() {
	phoneValidationQueue := make(chan model.PhoneValidationRequest, 100)
	go consumer.InitializePhoneValidationConsumer(phoneValidationQueue)

	router := router.Initialize(phoneValidationQueue)
	go router.Run()

	select {}
}
