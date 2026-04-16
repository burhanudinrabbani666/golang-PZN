package golangpzn

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Before Middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Middleware")

}

func (errorHandler *ErrorHandler) ServeHTTPError(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Terjadi Error")
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()

	errorHandler.ServeHTTPError(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Execute")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		panic("ups")

	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    localhost,
		Handler: errorHandler,
	}

	errorMiddleware := server.ListenAndServe()

	if errorMiddleware != nil {
		panic(errorMiddleware)
	}
}
