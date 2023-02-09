package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Alqo",
		MiddleName: "Mulky",
		LastName:   "Aria",
		Age:        26,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
