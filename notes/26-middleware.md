# Middleware

- Dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
- Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi

## Middleware di golang

- Sayangnya, di Go-Lang web tidak ada middleware
- Namun karena struktur handler yang baik menggunakan interface, kita bisa membuat middleware sendiri menggunakan handler

```go
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Before Middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Middleware")

}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Middleware")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    localhost,
		Handler: &logMiddleware,
	}

	errorMiddleware := server.ListenAndServe()

	if errorMiddleware != nil {
		panic(errorMiddleware)
	}
}
```

## Error Handler

- Kadang middleware juga biasa digunakan untuk melakukan error handler
- Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan mengubah panic tersebut menjadi error response
- Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan

Next: [Routing library](./27-routing-library.md)
