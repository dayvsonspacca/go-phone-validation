package model

type PhoneValidationRequest struct {
	Token                string
	Phone                string `json:"phone"`
	NationalIdentyNumber string `json:"national_identy_number"`
	CallbackUrl          string `json:"callback_url"`
}
