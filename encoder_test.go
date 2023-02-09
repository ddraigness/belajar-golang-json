package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Alqo",
		MiddleName: "Mulky",
		LastName:   "Aria",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
