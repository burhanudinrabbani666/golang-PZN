# Query parameter

- Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
- Query parameter biasanya digunakan untuk mengirim data dari client ke server
- Query parameter ditempatkan pada URL
- Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya

## url.URL

- Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
- Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang akan mengembalikan map

```go
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParams(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello/?name=bani", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
```

## Multiple Query Parameter

- Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
- Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
- Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya

```go
func MultipleQueryParams(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParams(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello/?first_name=Burhanudin&last_name=Rabbani", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
```

## Multiple Value Query Parameter

- Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
- Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
- Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda, misal :
- name=Eko&name=Kurniawan

```go
func MultipleParamsValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}

func TestMultipleQueryParamsValues(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello/?name=Burhanudin&name=Rabbani", nil)
	recorder := httptest.NewRecorder()

	MultipleParamsValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)j

}
```

Next: [Header](./10-header.md)
