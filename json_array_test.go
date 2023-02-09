package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Alqo",
		MiddleName: "Mulky",
		LastName:   "Aria",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Alqo","MiddleName":"Mulky","LastName":"Aria","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Alqo",
		Addresses: []Address{
			{
				Street:     "Jalan Nyasar",
				Country:    "Indonesia",
				PostalCode: "10650",
			},
			{
				Street:     "Jalan yang Lurus",
				Country:    "Indonesia",
				PostalCode: "10630",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}


func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Alqo","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Nyasar","Country":"Indonesia","PostalCode":"10650"},{"Street":"Jalan yang Lurus","Country":"Indonesia","PostalCode":"10630"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Nyasar","Country":"Indonesia","PostalCode":"10650"},{"Street":"Jalan yang Lurus","Country":"Indonesia","PostalCode":"10630"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}


func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Nyasar",
			Country:    "Indonesia",
			PostalCode: "10650",
		},
		{
			Street:     "Jalan yang Lurus",
			Country:    "Indonesia",
			PostalCode: "10630",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
	}
