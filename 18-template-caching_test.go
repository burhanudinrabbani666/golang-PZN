package golangpzn

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Global Vairable
//
//go:embed templates/*.gohtml
var Templates embed.FS
var MyTemplate = template.Must(template.ParseFS(Templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	MyTemplate.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Chacing")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
