package golangpzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writter, "Hello World")
}

func TestHttp(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, localhostFull, nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body) // Return []byte
	bodyString := string(body)

	fmt.Println(bodyString)

}
