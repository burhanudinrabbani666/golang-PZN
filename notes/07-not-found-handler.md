# Not Found handler

- Selain panic handler, Router juga memiliki not found handler
- Not found handler adalah handler yang dieksekusi ketika client mencoba melakukan request URL yang memang tidak terdapat di Router
- Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke http.NotFound, namun kita bisa mengubah nya
- Caranya dengan mengubah router.NotFound = http.Handler

```go
func TestNotFound(t *testing.T) {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halaman Tidak ditemukan")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Halaman Tidak ditemukan", string(body))
}
```

Next: [method not allowed handler](./08-method-not-allowed-handler.md)
