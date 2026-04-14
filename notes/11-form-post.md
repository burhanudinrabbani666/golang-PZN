# Form Post

## Gambaran Umum

| Konsep                  | Penjelasan                                                                          |
| ----------------------- | ----------------------------------------------------------------------------------- |
| Form GET                | Data form dikirim sebagai query parameter di URL                                    |
| Form POST               | Data form dikirim di dalam body HTTP request                                        |
| `r.ParseForm()`         | Wajib dipanggil sebelum membaca data form — mem-parsing body request                |
| `r.PostForm`            | Map berisi data form yang dikirim via POST, bertipe `url.Values`                    |
| `r.PostForm.Get("key")` | Mengambil satu nilai dari field form tertentu                                       |
| `Content-Type`          | Harus `application/x-www-form-urlencoded` agar Go bisa memparsing body sebagai form |

Saat HTML form disubmit dengan method `POST`, data tidak dikirim lewat URL melainkan lewat **body HTTP request**. Go menyediakan `r.PostForm` untuk membaca data tersebut, namun parsing harus dilakukan secara eksplisit terlebih dahulu dengan `r.ParseForm()`.

## Membaca Data Form POST

```go
func FormPost(w http.ResponseWriter, r *http.Request) {
    // Wajib dipanggil dulu sebelum mengakses r.PostForm
    // ParseForm() membaca body request dan mengubahnya menjadi map key-value
    // Jika body tidak bisa diparsing sebagai form, akan mengembalikan error
    err := r.ParseForm()
    if err != nil {
        panic(err)
    }

    // Setelah ParseForm(), data form bisa diakses via r.PostForm
    firstName := r.PostForm.Get("first_name")
    lastName := r.PostForm.Get("last_name")

    fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
    // Buat body request berisi data form dalam format key=value&key=value
    reqBody := strings.NewReader("first_name=Burhanudin&last_name=Rabbani")

    req := httptest.NewRequest("POST", "http://localhost:8080", reqBody)

    // Content-Type wajib diset agar Go tahu bahwa body adalah form data
    // Tanpa header ini, ParseForm() tidak akan membaca body dengan benar
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    recorder := httptest.NewRecorder()
    FormPost(recorder, req)

    res := recorder.Result()
    body, _ := io.ReadAll(res.Body)
    fmt.Println(string(body)) // Output: Hello, Burhanudin Rabbani
}
```

### Cara Kerja

```
Client mengirim HTTP POST request:

  POST /submit HTTP/1.1
  Host: localhost:8080
  Content-Type: application/x-www-form-urlencoded

  first_name=Burhanudin&last_name=Rabbani
  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
  Data form ada di body, bukan di URL

         ↓

r.ParseForm() membaca body dan mengubahnya:
  r.PostForm = url.Values{
      "first_name": ["Burhanudin"],
      "last_name":  ["Rabbani"],
  }

         ↓

r.PostForm.Get("first_name") → "Burhanudin"
r.PostForm.Get("last_name")  → "Rabbani"
```

## Perbedaan Form GET vs Form POST

|                     | Form GET                       | Form POST                         |
| ------------------- | ------------------------------ | --------------------------------- |
| Letak data          | Di URL sebagai query parameter | Di body HTTP request              |
| Contoh              | `/submit?name=bani&age=20`     | Body: `name=bani&age=20`          |
| Cara baca di Go     | `r.URL.Query().Get("key")`     | `r.PostForm.Get("key")`           |
| Wajib `ParseForm()` | Tidak                          | Ya                                |
| Cocok untuk         | Data kecil, pencarian, filter  | Data sensitif, form login, upload |
| Terlihat di URL     | Ya                             | Tidak                             |

## Alur Penggunaan

```
1. Client submit HTML form dengan method POST
2. Browser mengirim body: "first_name=Burhanudin&last_name=Rabbani"
3. Browser menyertakan header: Content-Type: application/x-www-form-urlencoded
4. Handler memanggil r.ParseForm() → body diparsing menjadi map
5. Data diakses via r.PostForm.Get("key")
6. Handler memproses data dan menulis response
```

## Catatan Penting

| Catatan                               | Detail                                                                                        |
| ------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ParseForm()` wajib dipanggil pertama | Mengakses `r.PostForm` tanpa `ParseForm()` akan menghasilkan map kosong                       |
| `Content-Type` harus tepat            | Harus `application/x-www-form-urlencoded` — bukan `application/json`                          |
| `r.PostForm` vs `r.Form`              | `r.PostForm` hanya membaca dari body POST; `r.Form` membaca dari body **dan** query parameter |
| Body hanya bisa dibaca sekali         | Setelah `ParseForm()` membaca body, body tidak bisa dibaca ulang                              |
| Nilai selalu bertipe `string`         | Sama seperti query parameter, semua nilai form adalah string                                  |

Next: [Response Code](./12-response-code.md)
