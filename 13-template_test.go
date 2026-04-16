package golangpzn

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	simpleTemplate := template.Must(template.New("SIMPLE").Parse(templateText))

	simpleTemplate.ExecuteTemplate(w, "SIMPLE", "Hello from Template")
}

func TestTemplateHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------//
// --------------------------------------------------------------------------------------------//

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	simpleTemplate := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	simpleTemplate.ExecuteTemplate(w, "simple.gohtml", "Hello from Template File")
}

func TestTemplateHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------//
// --------------------------------------------------------------------------------------------//

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	simpleTemplate := template.Must(template.ParseGlob("./templates/*.gohtml"))
	simpleTemplate.ExecuteTemplate(w, "hello.gohtml", "Hello from Template Directory")
}

func TestTemplateHtmlDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------//
// --------------------------------------------------------------------------------------------//


func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	simpleTemplate := template.Must(template.ParseFS(Templates, "templates/*.gohtml"))
	simpleTemplate.ExecuteTemplate(w, "simple.gohtml", "Hello from Template Embed")
}

func TestTemplateHtmlEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhostFull, nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
