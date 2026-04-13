# Form post

- Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method `GET` atau `POST`
- Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
- Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via **body HTTP request**
- Di Go-Lang, untuk mengambil data Form Post sangatlah mudah

## Request.PostForm

- Semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute `Request.PostForm`
- Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method `Request.ParseForm()` terlebih dahulu, method ini digunakan untuk melakukan parsing data body apakah bisa di parsing menjadi form data atau tidak, jika tidak bisa di parsing, maka akan menyebabkan error

```go
func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {

	reqBody := strings.NewReader("first_name=Burhanudin&last_name=Rabbani")
	req := httptest.NewRequest("POST", "http://localhost:8080", reqBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // Spesifikasi Standard web

	recorder := httptest.NewRecorder()
	FormPost(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
```

Next: [Response code](./12-response-code.md)
