package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Alqo")
	logJson(1)
	logJson(true)
	logJson([]string{"Alqo", "Mulky", "Aria"})
}