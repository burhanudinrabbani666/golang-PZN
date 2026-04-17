package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHttprouterParams(t *testing.T) {

	router := httprouter.New()

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Products 1", string(body))
}
