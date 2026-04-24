# JSON array

- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
- Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks (Object atau Array)
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice

```go
func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "Rabbani",
		Hobbies:    []string{"Games", "Coding", "Study"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Burhanudin","MiddleName":"D","LastName":"Rabbani","Hobbies":["Games","Coding","Study"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)

}
```

## Decode JSON Array

- Selain menggunakan Array pada attribute di Object
- Kita juga bisa melakukan encode atau decode langsung JSON Array nya
- Encode dan Decode JSON Array bisa menggunakan tipe data Slice

Next: [JSON tag](./06-json-tag.md)
