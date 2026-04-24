# encode JSON

- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function `json.Marshal(interface{})`
- Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal

```go
func LogJson(data any) {
	byte, errorJson := json.Marshal(data)

	if errorJson != nil {
		panic(errorJson)
	}

	// Byte harus di conversi ke string supaya bisa dibaca.
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {

	fmt.Println("Bani") // Kalau print string biasa tidak ada ""
	LogJson("Bani")     // "Bani", di Json string add ""
	LogJson(1)          // 1
	LogJson(true)       // true

	LogJson([]string{
		"Burhanudin",
		"D",
		"Rabbani",
	}) // ["Burhanudin", "D", "Rabbani"]
}
```

Next: [JSON object](./03-json-object.md)
