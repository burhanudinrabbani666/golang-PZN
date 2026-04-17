package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHttprouterPastterNameParams(t *testing.T) {

	router := httprouter.New()

	router.GET("/products/:id/items/:itemsId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemsId := p.ByName("itemsId")
		text := "Products: " + id + ", Item id: " + itemsId

		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Products: 1, Item id: 1", string(body))
}

func TestHttprouterPastterCatchALlParams(t *testing.T) {

	router := httprouter.New()

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image: " + image

		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image: /small/profile.png", string(body))
}
