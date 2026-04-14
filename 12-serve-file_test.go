package golangpzn

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notFound.html")
	}
}

func TestServeFileServer(t *testing.T) {

	server := http.Server{
		Addr:    localhost,
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOK string

//go:embed resources/notFound.html
var resourcesNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourcesOK)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {

	server := http.Server{
		Addr:    localhost,
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
