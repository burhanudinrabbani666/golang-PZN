package golangpzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAction(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/if.gohtml"))
	nameTemplate.ExecuteTemplate(w, "if.gohtml", map[string]any{
		"Title": "Template Action",
		// "Name":  "Burhanudin Rabbani",
	})

}

func TestTemplateAction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAction(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------- //
// --------------------------------------------------------------------------------------------- //

func TemplateComparator(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	nameTemplate.ExecuteTemplate(w, "comparator.gohtml", map[string]any{
		"Title":      "Template Action",
		"FinalValue": 50,
	})

}

func TestTemplateComparator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------- //
// --------------------------------------------------------------------------------------------- //

func TemplateRange(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/range.gohtml"))
	nameTemplate.ExecuteTemplate(w, "range.gohtml", map[string]any{
		"Title":   "Template Action Range",
		"Hobbies": []string{"Game", "Football", "Coding"},
	})

}

func TestTemplateRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// --------------------------------------------------------------------------------------------- //
// --------------------------------------------------------------------------------------------- //

func TemplateWith(w http.ResponseWriter, r *http.Request) {

	nameTemplate := template.Must(template.ParseFiles("./templates/with.gohtml"))
	nameTemplate.ExecuteTemplate(w, "with.gohtml", map[string]any{
		"Title": "Template Action WIth",
		"Name":  "Burhanudin Rabbani",
		"Address": map[string]any{
			"Street": "Raya",
			"City":   "Cirebon",
		},
	})

}

func TestTemplateWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
