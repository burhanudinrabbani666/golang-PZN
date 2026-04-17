# decode JSON

- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON
- Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan function json.Unmarshal(byte[], interface{})
- Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer

```go
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
```

Next: [JSON array](./05-json-array.md)
