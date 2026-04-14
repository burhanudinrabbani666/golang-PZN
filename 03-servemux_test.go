package golangpzn

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello World") })
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hi World") })
	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Images") })

	server := http.Server{
		Addr:    localhost,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestServerMuxUrlPattern(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello World") })
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hi World") })
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Images") })

	// Diprioritas kan dari pada hanya /images saja
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Thumbnails") })

	server := http.Server{
		Addr:    localhost,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
