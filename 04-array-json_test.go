package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "Rabbani",
		Hobbies:    []string{"Gaming", "Coding", "Reading"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
	/*
		Return:
		{
			"FirstName":"Burhanudin",
			"MiddleName":"D",
			"LastName":"Rabbani",
			"Hobbies":["Games","Coding","Study"],
		}
	*/

}

// -------------------------------------------- Decode -----------------------------------------
func TestJSONArrayDecode(t *testing.T) {

	// Simulasi JSON object dimana Keys sesuai dengan field Struct
	jsonString := `{
			"FirstName":"Burhanudin",
			"MiddleName":"D",
			"LastName":"Rabbani",
			"Hobbies":["Games","Coding","Study"]
			}`

	// Konversi ke array of byte
	jsonBytes := []byte(jsonString)

	// Buat Variable Kosong dengan type Customer
	// Untuk Menjadi wadah decode nya
	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)         // {Burhanudin D Rabbani [Games Coding Study]} <--- Slice biasa
	fmt.Println(customer.Hobbies) // [Games Coding Study] <----- Slice Biasa
}

func TestJSONArrayComplexEncode(t *testing.T) {

	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "Rabbani",
		Age:        23,
		IsMarried:  false,
		Hobbies:    []string{"Gaming", "Coding", "Reading"},
		Addresses: []Address{
			{
				Street:  "Jalan Belum Ada",
				Country: "Indonesia",
				Postal:  "12344",
			},
			{
				Street:  "Jalan lagi dibangun",
				Country: "Malaysia",
				Postal:  "23456",
			},
		},
	}

	bytes, errorCustomer := json.Marshal(customer)

	if errorCustomer != nil {
		panic(errorCustomer.Error())
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{
			"FirstName":"Burhanudin",
			"MiddleName":"D",
			"LastName":"Rabbani",
			"Hobbies":["Games","Coding","Study"]}`

	jsonBytes := []byte(jsonString)

	customer := Customer{}
	errorDecodeUser := json.Unmarshal(jsonBytes, &customer)

	if errorDecodeUser != nil {
		panic(errorDecodeUser)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}

// --------------------------------- Only Array -----------------------

func TestOnlyJSONArray(t *testing.T) {
	jsonString := `[
			{
				"Street":"Jalan Belum Ada",
				"Country":"Indonesia",
				"Postal":"12344"
			},
			{
				"Street":"Jalan lagi dibangun",
				"Country":"Malaysia",
				"Postal":"23456"
			}
		]`
	jsonBytes := []byte(jsonString)

	addresses := []Address{}

	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}

	for index, address := range addresses {
		fmt.Println("-------------------")
		fmt.Printf("Array ke: %d\n", index)
		fmt.Printf("Street: %s\n", address.Street)
		fmt.Printf("Country: %s\n", address.Country)
		fmt.Printf("Postal: %s\n", address.Postal)
	}
}
