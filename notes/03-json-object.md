# JSON object

- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
- Sedangkan value nya baru berupa

```json
	// Hanya kirim Value, tidak ideal atau real word case ❎
	"Bani"

	// Normal JSON ✅
  {
		"Name": "Bani"
	}
```

## Struct

- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct

```go
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "Rabbani",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	/*
		Return:
			{
				"FirstName":"Burhanudin",
				"MiddleName":"D",
				"LastName":"Rabbani",
			}
	*/

}
```

Next: [decode JSON](./04-decode-json.md)
