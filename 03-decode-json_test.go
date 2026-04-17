package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{
		"FirstName":"Burhanudin",
		"MiddleName":"D",
		"LastName":"Rabbani"
	}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}

	errorJson := json.Unmarshal(jsonBytes, customer)

	if errorJson != nil {
		panic(errorJson)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)

}
