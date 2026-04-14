# Request

## Gambaran Umum

| Field / Method | Tipe            | Penjelasan                                            |
| -------------- | --------------- | ----------------------------------------------------- |
| `r.Method`     | `string`        | HTTP method dari request (`GET`, `POST`, `PUT`, dst.) |
| `r.RequestURI` | `string`        | URL lengkap yang diminta, termasuk query string       |
| `r.URL`        | `*url.URL`      | URL yang sudah diparsing menjadi struct terstruktur   |
| `r.Header`     | `http.Header`   | Semua HTTP header yang dikirim client                 |
| `r.Body`       | `io.ReadCloser` | Isi body dari request (biasanya ada di POST/PUT)      |
| `r.RemoteAddr` | `string`        | Alamat IP dan port pengirim request                   |
| `r.Host`       | `string`        | Nama host yang dituju client                          |

`http.Request` adalah struct yang merepresentasikan satu HTTP request yang masuk dari browser atau client lainnya. Setiap kali ada request masuk, Go secara otomatis membungkus seluruh informasi request tersebut ke dalam struct ini dan meneruskannya ke handler sebagai parameter `r`.

## Membaca Informasi Request

```go
func TestRequest(t *testing.T) {

    var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
        // r adalah pointer ke http.Request yang berisi semua info request

        // Cetak HTTP method (GET, POST, PUT, DELETE, dst.)
        fmt.Fprintln(w, r.Method)

        // Cetak URI lengkap yang diminta client, contoh: /hello?name=bani
        fmt.Fprintln(w, r.RequestURI)
    }

    server := http.Server{
        Addr:    "localhost:8080",
        Handler: handler,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

### Cara Kerja

```
Browser mengirim:
  GET /hello?name=bani HTTP/1.1
  Host: localhost:8080

         ↓

Go membungkus request ke dalam http.Request:
  r.Method     = "GET"
  r.RequestURI = "/hello?name=bani"
  r.Host       = "localhost:8080"
  r.Header     = { "User-Agent": "...", ... }
  r.Body       = (kosong untuk GET)

         ↓

Handler menerima r dan membacanya sesuai kebutuhan
```

## Field-Field Penting pada Request

```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    // HTTP Method: GET, POST, PUT, DELETE, PATCH, dst.
    fmt.Fprintln(w, "Method     :", r.Method)

    // URI mentah yang dikirim client (termasuk query string)
    fmt.Fprintln(w, "RequestURI :", r.RequestURI)

    // URL yang sudah diparsing — bisa diakses per bagian
    fmt.Fprintln(w, "Path       :", r.URL.Path)
    fmt.Fprintln(w, "RawQuery   :", r.URL.RawQuery) // contoh: name=bani&age=20

    // Header yang dikirim client
    fmt.Fprintln(w, "User-Agent :", r.Header.Get("User-Agent"))

    // Alamat IP pengirim
    fmt.Fprintln(w, "RemoteAddr :", r.RemoteAddr)
}
```

### Perbedaan `RequestURI` vs `r.URL`

|                 | `r.RequestURI`                 | `r.URL`                           |
| --------------- | ------------------------------ | --------------------------------- |
| Tipe            | `string`                       | `*url.URL`                        |
| Isi             | URI mentah: `/hello?name=bani` | Struct terparsing                 |
| Akses query     | Harus diparse manual           | Langsung via `r.URL.Query()`      |
| Kapan digunakan | Hanya perlu string URI         | Perlu mengakses bagian-bagian URL |

## Alur Penggunaan

```
1. Client mengirim HTTP request ke server
2. Go menerima request dan membuat struct http.Request
3. Server meneruskan request ke Handler yang sesuai (via ServeMux atau langsung)
4. Handler membaca informasi dari r sesuai kebutuhan
5. Handler menulis response ke w (http.ResponseWriter)
6. Response dikirim balik ke client
```

## Catatan Penting

| Catatan                         | Detail                                                                              |
| ------------------------------- | ----------------------------------------------------------------------------------- |
| `r` selalu tersedia             | Setiap handler selalu menerima `r *http.Request` secara otomatis dari Go            |
| `r.Body` perlu ditutup          | Jika membaca body, panggil `defer r.Body.Close()` agar tidak bocor memori           |
| `r.URL` vs `r.RequestURI`       | Gunakan `r.URL` jika ingin mengakses path atau query secara terstruktur             |
| `r.Method` selalu huruf kapital | Go mengembalikan method dalam huruf kapital: `"GET"`, bukan `"get"`                 |
| Request bersifat read-only      | Jangan memodifikasi struct request — gunakan variabel lokal jika perlu transformasi |

Next: [Http Test](./8-http-test.md)
