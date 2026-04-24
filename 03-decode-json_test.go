package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {

	// Simulasi JSON object
	jsonString := `{
		"FirstName":"Burhanudin",
		"MiddleName":"D",
		"LastName":"Rabbani"
	}`

	// Ubah String menjadi string byte
	// Ingat return dari json.marshal() adalah bytes atau []byte
	// lalu kita conversi dengan cara string(bytes)
	// Jadi secara logic kita hanya perlu membaliknya
	jsonBytes := []byte(jsonString)
	customer := Customer{}

	// Kita masukan langsug ke customer dengan Pointer. pass by data
	errorJson := json.Unmarshal(jsonBytes, &customer)

	if errorJson != nil {
		panic(errorJson)
	}

	// Sekarang customer terisi dengan json yang sudah diencode
	fmt.Println(customer)            // {Burhanudin D Rabbani }
	fmt.Println(customer.FirstName)  // Burhanudin
	fmt.Println(customer.MiddleName) // D
	fmt.Println(customer.LastName)   // Rabbai

	// Ingat ⚠️ JSON Itu dengan "",  Jadi Kalau value tidak ada "" setelah di ubah dari JSON. berarti sudah bukan lagi type data JSON
}
