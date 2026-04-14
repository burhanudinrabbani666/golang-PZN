# HTTP Test

## Gambaran Umum

| Komponen                    | Penjelasan                                                            |
| --------------------------- | --------------------------------------------------------------------- |
| `httptest.NewRequest()`     | Membuat simulasi HTTP request untuk keperluan testing                 |
| `httptest.NewRecorder()`    | Membuat perekam response yang menggantikan `http.ResponseWriter` asli |
| `recorder.Result()`         | Mengambil HTTP response yang sudah direkam setelah handler dijalankan |
| `io.ReadAll(response.Body)` | Membaca isi body response menjadi `[]byte`                            |

Package `net/http/httptest` memungkinkan kita menguji handler web **tanpa harus menjalankan server HTTP sungguhan**. Handler dipanggil langsung sebagai fungsi biasa, sehingga pengujian lebih cepat, terisolasi, dan tidak bergantung pada port atau jaringan.

## Contoh Penggunaan

```go
// Handler yang ingin kita uji
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintln(writer, "Hello World")
}

func TestHttp(t *testing.T) {
    // 1. Buat simulasi request: method GET ke URL tersebut, tanpa body (nil)
    request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

    // 2. Buat recorder sebagai pengganti http.ResponseWriter
    //    Semua output yang ditulis handler akan tersimpan di sini
    recorder := httptest.NewRecorder()

    // 3. Panggil handler langsung — tidak perlu menjalankan server
    HelloHandler(recorder, request)

    // 4. Ambil hasil response dari recorder
    response := recorder.Result()

    // 5. Baca isi body response
    body, _ := io.ReadAll(response.Body)
    bodyString := string(body)

    fmt.Println(bodyString) // Output: Hello World
}
```

### Cara Kerja

```
Testing Biasa (dengan server):              Testing dengan httptest:
  Jalankan server        ──────────────►   Tidak perlu server
  Kirim HTTP request     ──────────────►   httptest.NewRequest()
  Tunggu response        ──────────────►   httptest.NewRecorder()
  Baca response          ──────────────►   recorder.Result()
  Matikan server         ──────────────►   Tidak perlu
```

```
httptest.NewRequest()  ──► request  ──┐
                                      ├──► HelloHandler(recorder, request)
httptest.NewRecorder() ──► recorder ──┘
                                      │
                                      ▼
                              recorder.Result()  ──► *http.Response
                                      │
                                      ▼
                              io.ReadAll(response.Body) ──► "Hello World\n"
```

## `httptest.NewRequest()`

```go
request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

// Menambahkan header ke request simulasi
request.Header.Set("Authorization", "Bearer token123")
request.Header.Set("Content-Type", "application/json")
```

| Parameter | Tipe        | Contoh                                               |
| --------- | ----------- | ---------------------------------------------------- |
| `method`  | `string`    | `http.MethodGet`, `http.MethodPost`, `"DELETE"`      |
| `url`     | `string`    | `"http://localhost:8080/hello?name=bani"`            |
| `body`    | `io.Reader` | `nil` untuk GET, `strings.NewReader(...)` untuk POST |

## `httptest.NewRecorder()`

```go
recorder := httptest.NewRecorder()

// Setelah handler dipanggil, kita bisa mengakses:
response := recorder.Result()

fmt.Println(response.StatusCode)          // contoh: 200
fmt.Println(response.Header.Get("Content-Type")) // header response
body, _ := io.ReadAll(response.Body)
fmt.Println(string(body))                 // isi body response
```

`ResponseRecorder` merekam semua yang ditulis handler ke `w` — termasuk status code, header, dan body — lalu menyimpannya agar bisa kita periksa dalam assertion test.

## Alur Penggunaan

```
1. Tulis handler function yang ingin diuji
2. Buat request simulasi  → httptest.NewRequest(method, url, body)
3. Buat recorder          → httptest.NewRecorder()
4. Panggil handler        → MyHandler(recorder, request)
5. Ambil response         → recorder.Result()
6. Periksa hasilnya       → status code, header, body
```

## Catatan Penting

| Catatan                               | Detail                                                                         |
| ------------------------------------- | ------------------------------------------------------------------------------ |
| Tidak perlu server                    | Handler dipanggil langsung sebagai fungsi — tidak ada koneksi jaringan         |
| `recorder` = `ResponseWriter`         | `httptest.NewRecorder()` mengimplementasikan interface `http.ResponseWriter`   |
| Body perlu dibaca dengan `io.ReadAll` | `response.Body` bertipe `io.ReadCloser`, perlu dibaca secara eksplisit         |
| URL bisa pakai path saja              | `httptest.NewRequest` menerima `"/hello"` maupun `"http://localhost/hello"`    |
| Cocok digabung dengan `testing.T`     | Gunakan `t.Errorf` atau `assert` dari Testify untuk memvalidasi hasil response |

Next: [Query Parameter](./9-query-parameter.md)
