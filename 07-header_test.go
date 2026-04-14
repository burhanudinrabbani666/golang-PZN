package golangpzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)

}

func TestReqHeader(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, localhostFull, nil)
	request.Header.Add("Content-Type", "application/json") // Set Header with Key "Content-Type"

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Burhanudin Rabbani")
	fmt.Fprint(w, "OK")
}

func TestResHeader(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, localhostFull, nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))

}
