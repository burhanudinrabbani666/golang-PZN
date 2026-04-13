# Template action

- Go-Lang template mendukung perintah action, seperti percabangan, perulangan dan lain-lain

## If Else

- `{{if .Value}} T1 {{end}}`, jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong, tidak ada yang dieksekusi
- `{{if .Value}} T1 {{else}} T2 {{end}}`, jika value tidak kosong, maka T1 akan dieksekusi, jika kosong, T2 yang akan dieksekusi
- `{{if .Value1}} T1 {{else if .Value2}} T2 {{else}} T3 {{end}}`, jika Value1 tidak kosong, maka T1 akan dieksekusi, jika Value2 tidak kosong, maka T2 akan dieksekusi, jika tidak semuanya, maka T3 akan dieksekusi

```html
<body>
  {{if.Name}}
  <h1>Hello {{.Name}}</h1>
  {{else}}
  <h1>Hello</h1>
  {{end}}
</body>
```

## Operator Perbandingan

Go-Lang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan perbandingan number di if statement, berikut adalah operator nya :

- eq artinya arg1 == arg2
- ne artinya arg1 != arg2
- lt artinya arg1 < arg2
- le artinya arg1 <= arg2
- gt artinya arg1 > arg2
- ge artinya arg1 >= arg2

## Kenapa Operatornya di Depan?

- Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
- Jadi saat kita menggunakan {{eq First Second}}, sebenarnya dia akan memanggil function eq dengan parameter First dan Second : eq(First, Second)

```go
func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", Page{
		Title:      "Template Action Operator",
		Name:       "Bani",
		FinalValue: 50,
	})
}

func TestTemplateDataOparator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
```

```html
<body>
  {{if ge .FinalValue 80}}
  <h1>Good</h1>
  {{else if ge .FinalValue 60}}
  <h1>Nice Try</h1>
  {{else}}
  <h1>Try Again</h1>
  {{end}}
</body>
```

## Range

- Range digunakan untuk melakukan iterasi data template
- Tidak ada perulangan biasa seperti menggunakan for di Go-Lang template
- mYang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map atau channel
- {{range $index, $element := .Value}} T1 {{end}}, jika value memiliki data, maka T1 akan dieksekusi sebanyak element value, dan kita bisa menggunakan $index untuk mengakses index dan $element untuk mengakses element
- {{range $index, $element := .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value tidak memiliki element apapun, maka T2 yang akan dieksekusi

```go
func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]any{
		"Title":   "Template Action range",
		"Hobbies": []string{"Game", "Read", "Coding"},
	})
}
```

```html
<body>
  {{range $index, $element := .Hobbies}}
  <h1>{{$element}}</h1>
  {{else}}
  <h1>Anda Tidak punya Hobi</h1>
  {{end}}
</body>
```

## With

- Kadang kita sering membuat nested struct
- Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue
- Di template terdapat action with, yang bisa digunakan mengubah scope dot menjadi object yang kita mau
- {{with .Value}} T1 {{end}}, jika value tidak kosong, di T1 semua dot akan merefer ke value
- {{with .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value kosong, maka T2 yang akan dieksekusi

```go
func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]any{
		"Title": "Template Action range",
		"Name":  "Bani",
		"Address": map[string]any{
			"Street": "Jalan Belum ada",
			"City":   "Cirebon",
		},
	})
}

```

```html
<h1>Name: {{.Name}}</h1>
{{with .Address}}
<h1>Address Street: {{.Street}}</h1>
<h1>Address City: {{.City}}</h1>
{{else}}
<h1>Anda Tidak Punya alamat</h1>
{{end}}
```

## Comment

- Template juga mendukung komentar
- Komentar secara otomatis akan hilang ketika template text di parsing
- Untuk membuat komentar sangat sederhana, kita bisa gunakan {{/* Contoh Komentar */}}

Next: [Template layout](./19-template-layout.md)
