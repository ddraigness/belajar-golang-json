package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONTag(t *testing.T) {
	product := Product{
		Id: "P0001",
		Name: "Apple MacBook Pro",
		ImageURL: "http://flashklik.com/product/1",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple MacBook Pro","price":0,"Image_URL":"http://flashklik.com/product/1"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(product)
}