package main

import (
	"github.com/dayvsonspacca/go-phone-validation/consumer"
	"github.com/dayvsonspacca/go-phone-validation/database"
	"github.com/dayvsonspacca/go-phone-validation/handler"
	"github.com/dayvsonspacca/go-phone-validation/router"
)

func main() {
	database.Initialize()

	phoneValidationQueue := make(chan handler.PhoneValidationData, 100)

	go consumer.InitializePhoneValidationConsumer(phoneValidationQueue)

	router.Initialize(phoneValidationQueue).Run()
}
