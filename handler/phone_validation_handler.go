package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dayvsonspacca/go-phone-validation/request"
)

type PhoneValidationResult string

const (
	MATCHED   PhoneValidationResult = "MATCHED"
	UNMATCHED PhoneValidationResult = "UNMATCHED"
	NO_DATA   PhoneValidationResult = "NO_DATA"
)

type PhoneValidationResponse struct {
	Token  string                `json:"token"`
	Result PhoneValidationResult `json:"result"`
}

type PhoneValidationData struct {
	Token                  string
	PhoneValidationRequest request.PhoneValidationRequest
}

func HandlerPhoneValidation(phoneValidationData PhoneValidationData) {
	response := PhoneValidationResponse{
		Token:  phoneValidationData.Token,
		Result: validatePhoneNumber(phoneValidationData.PhoneValidationRequest.Phone, phoneValidationData.PhoneValidationRequest.NationalIdentyNumber),
	}

	sendPhoneValidationResponse(phoneValidationData.PhoneValidationRequest.CallbackUrl, response)
}

func validatePhoneNumber(phone string, nationalIdentyNumber string) PhoneValidationResult {
	// Access database, check if exists an entry with same number phone and person identy number.

	return MATCHED
}

func sendPhoneValidationResponse(callback_url string, response PhoneValidationResponse) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to parse JSON response: %s", err.Error())
		return
	}

	resp, err := http.Post(callback_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Failed to send callback response: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response sent to: %s Result: %s\n", callback_url, response.Result)
}
