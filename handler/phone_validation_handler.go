package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dayvsonspacca/go-phone-validation/database"
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
	fmt.Printf("Handling PhoneValidationData with token: %s\n", phoneValidationData.Token)

	response := PhoneValidationResponse{
		Token:  phoneValidationData.Token,
		Result: validatePhoneNumber(phoneValidationData.PhoneValidationRequest.Phone, phoneValidationData.PhoneValidationRequest.NationalIdentyNumber),
	}

	sendPhoneValidationResponse(phoneValidationData.PhoneValidationRequest.CallbackUrl, response)
}

func validatePhoneNumber(phone string, nationalIdentyNumber string) PhoneValidationResult {
	db := database.GetConnection()

	var status string
	err := db.QueryRow(
		"SELECT status FROM validated_phones WHERE phone = ? AND national_identy_number = ? LIMIT 1", phone, nationalIdentyNumber,
	).Scan(&status)

	if err != nil {
		if err == sql.ErrNoRows {
			return NO_DATA
		}

		return NO_DATA // [ TODO ] - Maybe not return NO_DATA when other erro than ErrNoRows occur
	}

	return PhoneValidationResult(status)
}

func sendPhoneValidationResponse(callback_url string, response PhoneValidationResponse) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to parse JSON response: %s", err.Error())
		return
	}

	resp, err := http.Post(callback_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Failed to send callback response: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response sent to: %s Result: %s\n", callback_url, response.Result)
}
