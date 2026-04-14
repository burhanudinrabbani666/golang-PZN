# Query Parameter

## Gambaran Umum

| Konsep            | Penjelasan                                                                  |
| ----------------- | --------------------------------------------------------------------------- |
| Query parameter   | Data yang dikirim client lewat URL, setelah tanda `?`                       |
| `r.URL.Query()`   | Mengambil semua query parameter sebagai `map[string][]string`               |
| `.Get("key")`     | Mengambil satu nilai dari key tertentu — jika tidak ada, mengembalikan `""` |
| `query["key"]`    | Mengambil semua nilai dari satu key sebagai `[]string`                      |
| Pemisah parameter | `?` untuk parameter pertama, `&` untuk parameter berikutnya                 |

Query parameter adalah cara umum untuk mengirim data dari client ke server melalui URL. Data ini terlihat langsung di URL sehingga cocok untuk pencarian, filter, atau data yang boleh dibagikan lewat link.

## Query Parameter Tunggal

```
Format URL: /hello?name=bani
                   ^^^^  ^^^^
                   key   value
```

```go
func SayHello(w http.ResponseWriter, r *http.Request) {
    // Query() mengembalikan map[string][]string dari semua query parameter
    // Get("name") mengambil nilai pertama dari key "name"
    // Jika key tidak ada, Get() mengembalikan string kosong ""
    name := r.URL.Query().Get("name")

    if name == "" {
        fmt.Fprint(w, "Hello")
    } else {
        fmt.Fprintf(w, "Hello %s", name)
    }
}

func TestQueryParams(t *testing.T) {
    // Simulasi request dengan query parameter name=bani
    request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=bani", nil)
    recorder := httptest.NewRecorder()

    SayHello(recorder, request)

    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body)) // Output: Hello bani
}
```

### Cara Kerja

```
URL: /hello?name=bani

r.URL.Query() menghasilkan:
  map[string][]string{
      "name": ["bani"],
  }

.Get("name") → "bani"
.Get("age")  → "" (key tidak ada)
```

## Multiple Query Parameter

Gunakan `&` untuk memisahkan beberapa query parameter dalam satu URL.

```
Format URL: /hello?first_name=Burhanudin&last_name=Rabbani
                   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                   param1                param2
```

```go
func MultipleQueryParams(w http.ResponseWriter, r *http.Request) {
    // Setiap key diakses secara terpisah menggunakan Get()
    firstName := r.URL.Query().Get("first_name")
    lastName := r.URL.Query().Get("last_name")

    fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParams(t *testing.T) {
    // Dua query parameter dipisahkan dengan &
    request := httptest.NewRequest(
        http.MethodGet,
        "http://localhost:8080/hello?first_name=Burhanudin&last_name=Rabbani",
        nil,
    )
    recorder := httptest.NewRecorder()

    MultipleQueryParams(recorder, request)

    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body)) // Output: Hello Burhanudin Rabbani
}
```

### Cara Kerja

```
URL: /hello?first_name=Burhanudin&last_name=Rabbani

r.URL.Query() menghasilkan:
  map[string][]string{
      "first_name": ["Burhanudin"],
      "last_name":  ["Rabbani"],
  }
```

## Multiple Value pada Satu Key

Satu key query parameter bisa memiliki beberapa nilai sekaligus, cukup ulangi key yang sama dengan value berbeda.

```
Format URL: /hello?name=Burhanudin&name=Rabbani
                   ^^^^^^^^^^^^^   ^^^^^^^^^^^^
                   key=value1      key=value2 (key sama, value berbeda)
```

```go
func MultipleParamValues(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    // Akses semua nilai dari key "name" sekaligus sebagai []string
    // Gunakan query["name"] bukan .Get("name") agar semua nilai terbaca
    names := query["name"]

    fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}

func TestMultipleQueryParamValues(t *testing.T) {
    // Key "name" muncul dua kali dengan nilai berbeda
    request := httptest.NewRequest(
        http.MethodGet,
        "http://localhost:8080/hello?name=Burhanudin&name=Rabbani",
        nil,
    )
    recorder := httptest.NewRecorder()

    MultipleParamValues(recorder, request)

    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body)) // Output: Hello Burhanudin Rabbani
}
```

### Cara Kerja

```
URL: /hello?name=Burhanudin&name=Rabbani

r.URL.Query() menghasilkan:
  map[string][]string{
      "name": ["Burhanudin", "Rabbani"],  ← semua value tersimpan dalam slice
  }

.Get("name")    → "Burhanudin"  (hanya nilai pertama)
query["name"]   → ["Burhanudin", "Rabbani"]  (semua nilai)
```

## Alur Penggunaan

```
1. Client mengirim URL dengan query parameter → /hello?name=bani&age=20
2. Go memparsing URL → r.URL bertipe *url.URL
3. r.URL.Query() mengubah query string → map[string][]string
4. .Get("key")    → ambil satu nilai (string)
5. query["key"]   → ambil semua nilai ([]string)
6. Handler memproses nilai dan menulis response
```

## Catatan Penting

| Catatan                            | Detail                                                                                      |
| ---------------------------------- | ------------------------------------------------------------------------------------------- |
| `.Get()` hanya ambil nilai pertama | Jika key punya banyak nilai, gunakan `query["key"]` untuk mengambil semua                   |
| Key tidak ada → string kosong      | `r.URL.Query().Get("tidakada")` mengembalikan `""`, bukan error                             |
| Tipe data selalu `string`          | Semua nilai query parameter adalah string — konversi manual jika butuh `int` atau tipe lain |
| Query parameter terlihat di URL    | Jangan gunakan untuk data sensitif seperti password atau token                              |
| Nama key bersifat case-sensitive   | `"name"` dan `"Name"` dianggap key yang berbeda                                             |

Next: [Header](./10-header.md)
