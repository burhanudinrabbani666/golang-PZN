# Template function

- Selain mengakses field, dalam template, function juga bisa diakses
- Cara mengakses function sama seperti mengakses field, namun jika function tersebut memiliki parameter, kita bisa gunakan tambahkan parameter ketika memanggil function di template nya
- `{{.FunctionName}}` memanggil field FunctionName atau function `FunctionName()`
- `{{.FunctionName “eko”, “kurniawan”}}` memanggil function `FunctionName(“eko”, “kurniawan”)`

```go
type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

func TemplateFunction(writter http.ResponseWriter, request *http.Request) {

	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Aisa"}}`))
	t.ExecuteTemplate(writter, "FUNCTION", MyPage{Name: "Bani"})

}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
```

## Global Function

- Go-Lang template memiliki beberapa global function
- Global function adalah function yang bisa digunakan secara langsung, tanpa menggunakan template data
- Berikut adalah beberapa global function di Go-Lang template
- https://github.com/golang/go/blob/master/src/text/template/funcs.go

```go
func TemplateFunctionGlobal(writter http.ResponseWriter, request *http.Request) {

	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writter, "FUNCTION", MyPage{Name: "Bani"})

}
```

## Menambah Global Function

- Kita juga bisa menambah global function
- Untuk menambah global function, kita bisa menggunakan method Funcs pada template
- Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum melakukan parsing template

```go
func TemplateFunctionCreateGlobal(writter http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")

	// Register Funcs
	t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse("{{upper .Name}}"))

	t.ExecuteTemplate(writter, "FUNCTION", MyPage{
		Name: "Bani",
	})

}
```

## Function Pipelines

- Go-Lang template mendukung function pipelines, artinya hasil dari function bisa dikirim ke function berikutnya
- Untuk menggunakan function pipelines, kita bisa menggunakan tanda | , misal :
- {{ sayHello .Name | upper }}, artinya akan memanggil global function sayHello(Name) hasil dari sayHello(Name) akan dikirim ke function upper(hasil)
- Kita bisa menambahkan function pipelines lebih dari satu

```go
func TemplateFunctionPipeline(writter http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")

	// Register Funcs
	t.Funcs(map[string]any{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse("{{sayHello .Name | upper}}"))

	t.ExecuteTemplate(writter, "FUNCTION", MyPage{
		Name: "burhanudin rabbani",
	})

}
```

Next: [Template caching](./21-template-caching.md)
