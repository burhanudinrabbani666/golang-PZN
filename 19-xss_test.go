package golangpzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TemplateAutoEscape(writer http.ResponseWriter, response *http.Request){
	MyTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "XSS Test",
		"Body": template.HTML("<p>Ini Adalah Auto Escape</p><script>alert('Anda di Hack')</script>"),

	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}


func TestTemplateAutoEscapeSever(t *testing.T) {
	server := http.Server{
		Addr: localhost,
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()

	if err != nil{
		panic(err)
	}
}


// --------------------------------------------------------- //
// --------------------------------------------------------- //

func TemplateXSS(writer http.ResponseWriter, request *http.Request){
	MyTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "XSS Test",
		"Body": template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost+"/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}


func TestTemplateXssServer(t *testing.T) {
	server := http.Server{
		Addr: localhost,
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()

	if err != nil{
		panic(err)
	}
}
