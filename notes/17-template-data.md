# Template data

- Saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis
- Hal ini bisa kita lakukan dengan cara menggunakan data struct atau map
- Namun perlu dilakukan perubahan di dalam text template nya, kita perlu memberi tahu Field atau Key mana yang akan kita gunakan untuk mengisi data dinamis di template
- Kita bisa menyebutkan dengan cara seperti ini {{.NamaField}}

```go
func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]any{
		"Title": "Template Data Map",
		"Name":  "Burhanudin Rabbani",
		"Address": map[string]any{
			"Street": "Cirebon Map",
		},
	})
}

func TestDataTemplateMap(t *testing.T) {
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
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Struct",
		Name:  "Burhanudin Rabbani",
		Address: AddressDetail{
			City: "Cirebon",
		},
	})
}

func TestDataTemplatStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
```

Next: [Template action](./18-template-action.md)
