package golangpzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/name.gohtml"))
	nameTemplate.ExecuteTemplate(w, "name.gohtml", map[string]any{
		"Title": "Template Data",
		"Name":  "Burhanudin Rabbani",
		"Address": map[string]any{
			"City": "Cirebon",
		},
	})

}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

type AddressDetail struct {
	City string
}

type Page struct {
	Title   string
	Name    string
	Address AddressDetail
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/name.gohtml"))
	nameTemplate.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template data Struct",
		Name:  "Burhanudin Rabbani",
		Address: AddressDetail{
			City: "Cirebon",
		},
	})
}

func TestTemplateStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
