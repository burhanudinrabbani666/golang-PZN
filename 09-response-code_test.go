package golangpzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name Is Empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func TestResponceCodeInvalid(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull, nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(recorder.Body)

	fmt.Println(response.Status)
	fmt.Println(string(body))

}

func TestResponceCodeValid(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull+"/?name=Burhanudin", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(recorder.Body)

	fmt.Println(response.Status)
	fmt.Println(string(body))

}
