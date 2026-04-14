package golangpzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseGlob("./templates/layout/*.gohtml"))

	nameTemplate.ExecuteTemplate(w, "body", map[string]any{
		"Title": "Template Layout",
		"Name":  "Burhanudin Rabbani",
	})

}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
