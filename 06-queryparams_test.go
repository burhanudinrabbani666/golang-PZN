package golangpzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Sayhello(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func TestQueryParams(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull+"/?name=bani", nil)
	recorder := httptest.NewRecorder()

	Sayhello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleQueryParams(w http.ResponseWriter, r *http.Request) {

	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)

}

func TestMultipleQueryParams(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull+"/?first_name=burhanudin&last_name=rabbani", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleQueryParamsValue(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	name := query["name"]
	fmt.Fprintf(w, "Hello %s", strings.Join(name, " "))

}

func TestMultipleQueryParamsValue(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull+"/?name=burhanudin&name=rabbani", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParamsValue(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
