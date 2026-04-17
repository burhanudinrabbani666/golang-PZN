# encode JSON

- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function json.Marshal(interface{})
- Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal

```go
func LogJson(data any) {
	byte, errorJson := json.Marshal(data)

	if errorJson != nil {
		panic(errorJson)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJson("Bani")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Burhanudin", "D", "Rabbani"})
}
```

Next: [JSON object](./03-json-object.md)
