# Method not allowed handler

- Saat menggunakan ServeMux, kita tidak bisa menentukan HTTP Method apa yang digunakan untuk Handler
- Namun pada Router, kita bisa menentukan HTTP Method yang ingin kita gunakan, lantas apa yang terjadi jika client tidak mengirim HTTP Method sesuai dengan yang kita tentukan?
- Maka akan terjadi error Method Not Allowed
- Secara default, jika terjadi error seperti ini, maka Router akan memanggil function http.Error
- Jika kita ingin mengubahnya, kita bisa gunakan router.MethodNotAllowed = http.Handler

```go
func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusMethodNotAllowed)
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")

	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "405", string(body))
}
```

Next: [Middleware](./09-middleware.md)
