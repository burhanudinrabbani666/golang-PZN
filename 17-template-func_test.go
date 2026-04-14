package golangpzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type User struct {
	Name string
}

func (user User) SayHello(name string) string {
	return "Hello " + name + ", My name is " + user.Name
}

func TemplateFunction(writter http.ResponseWriter, request *http.Request) {

	sayHelloTemplate := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Aisa"}}`)) // Automatic
	sayHelloTemplate.ExecuteTemplate(writter, "FUNCTION", User{Name: "Bani"})

}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writter http.ResponseWriter, request *http.Request) {

	sayHelloTemplate := template.Must(template.New("FUNCTION").Parse(`{{len "Aisa"}}`)) // Global Function
	sayHelloTemplate.ExecuteTemplate(writter, "FUNCTION", User{Name: "Bani"})

}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// ------------------------------------------------------------------------------------------- //
// ------------------------------------------------------------------------------------------- //

func TemplateFunctionCreateGlobal(writter http.ResponseWriter, request *http.Request) {

	functionTemplate := template.New("FUNCTION")

	functionTemplate = functionTemplate.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	templateToRender := template.Must(functionTemplate.Parse("{{upper .Name}}"))

	templateToRender.ExecuteTemplate(writter, "FUNCTION", map[string]any{
		"Name": "Burhanudin Rabbani",
	})

}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// ------------------------------------------------------------------------------------------- //
// ------------------------------------------------------------------------------------------- //

func TemplateFunctionPipeLine(writter http.ResponseWriter, request *http.Request) {

	functionTemplate := template.New("FUNCTION")

	functionTemplate = functionTemplate.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"sayHello": func(value string) string {
			return "Hello " + value
		},
	})

	templateToRender := template.Must(functionTemplate.Parse("{{sayHello .Name | upper}}"))

	templateToRender.ExecuteTemplate(writter, "FUNCTION", map[string]any{
		"Name": "Burhanudin Rabbani",
	})

}

func TestTemplateFunctionPipeLine(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeLine(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
