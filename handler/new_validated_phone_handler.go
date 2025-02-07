package handler

import (
	"fmt"

	"github.com/dayvsonspacca/go-phone-validation/database"
	"github.com/dayvsonspacca/go-phone-validation/request"
)

func HandlerNewValidatedPhoneNumber(newPhoneValidatedRequest request.NewPhoneValidatedRequest) int64 {
	db := database.GetConnection()
	if db == nil {
		fmt.Println("Database connection is nil")
		return 0
	}

	query := `INSERT INTO validated_phones (phone, national_identy_number, status) VALUES (?, ?, ?)`
	fmt.Printf("Executing query: %s with values: %s, %s %s\n",
		query, newPhoneValidatedRequest.Phone, newPhoneValidatedRequest.NationalIdentyNumber, newPhoneValidatedRequest.Status)

	row, err := db.Exec(query, newPhoneValidatedRequest.Phone, newPhoneValidatedRequest.NationalIdentyNumber, newPhoneValidatedRequest.Status)
	if err != nil {
		fmt.Printf("Failed to insert validated phone: %v\n", err)
		return 0
	}

	id, err := row.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get last inserted ID: %v\n", err)
		return 0
	}

	return id
}
