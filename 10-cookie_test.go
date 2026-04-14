package golangpzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := new(http.Cookie)
	cookie.Name = "X-Bani-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Succes Creted Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("X-Bani-Name")

	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprint(w, cookie.Raw)
		fmt.Fprintf(w, "Hello %s", name)
	}

}

// ---------------------- Test --------------------------------- //
func TestCookie(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    localhost,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhostFull+"/?name=Burhanudin", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {

	cookie := new(http.Cookie)
	cookie.Name = "X-Bani-Name"
	cookie.Value = "Burhanudin Rabbani Get Cookie"

	request := httptest.NewRequest(http.MethodGet, localhostFull, nil)
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
