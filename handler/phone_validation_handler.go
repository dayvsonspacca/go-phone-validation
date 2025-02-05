package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func ProcessPhoneValidationRequest(phone string, nationalIdentyNumer string) PhoneValidationResult {
	// Access database, check if exists an entry with same number phone and person identy number.

	return MATCHED
}

func SendPhoneValidationResponse(callback_url string, response PhoneValidationResponse) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(callback_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response sent to: %s Result: %s\n", callback_url, response.Result)
}
