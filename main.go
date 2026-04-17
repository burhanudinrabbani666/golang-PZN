package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	// ------------------ Server ------------------- //
	// ------------------------- ------------------- //

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	serverError := server.ListenAndServe()

	if serverError != nil {
		panic(serverError)
	}
}
