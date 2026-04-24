package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street  string
	Country string
	Postal  string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	IsMarried  bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {

	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "Rabbani",
		Age:        23,
		IsMarried:  false,
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
	/*
		Return:
			{
				"FirstName":"Burhanudin",
				"MiddleName":"D",
				"LastName":"Rabbani",
				"Age":23,
				"IsMarried":false,
				"Hobbies":null,
				"Addresses":null
			}
	*/

}
