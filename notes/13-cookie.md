# Cookie

## Gambaran Umum

| Konsep                        | Penjelasan                                                                                     |
| ----------------------------- | ---------------------------------------------------------------------------------------------- |
| Stateless                     | HTTP tidak menyimpan state antara request — setiap request dianggap baru                       |
| Cookie                        | Data key-value yang dibuat server, disimpan di browser, dan dikirim otomatis di setiap request |
| `http.SetCookie()`            | Mengirim cookie dari server ke browser melalui response header                                 |
| `r.Cookie("name")`            | Membaca satu cookie dari request yang dikirim client                                           |
| `recorder.Result().Cookies()` | Mengambil semua cookie dari response dalam unit test                                           |
| `req.AddCookie()`             | Menambahkan cookie ke request simulasi dalam unit test                                         |

HTTP bersifat **stateless** — server tidak mengingat siapa yang mengirim request sebelumnya. Cookie adalah solusi standar untuk menyimpan informasi sesi di sisi browser, sehingga server bisa mengenali client yang sama di request berikutnya.

## Cara Kerja Cookie

```
Request pertama (login):
  Client → Server: POST /login
  Server → Client: Set-Cookie: X-Bani-Name=Burhanudin
                   ↑ server menyuruh browser menyimpan cookie ini

Request berikutnya:
  Client → Server: GET /dashboard
                   Cookie: X-Bani-Name=Burhanudin
                   ↑ browser otomatis menyertakan cookie di setiap request

  Server membaca cookie → tahu siapa clientnya → tidak perlu login lagi
```

## Membuat dan Membaca Cookie

```go
// Handler untuk membuat cookie baru di browser client
func SetCookie(w http.ResponseWriter, r *http.Request) {
    cookie := new(http.Cookie)
    cookie.Name = "X-Bani-Name"                  // Nama/key cookie
    cookie.Value = r.URL.Query().Get("name")      // Nilai diambil dari query parameter
    cookie.Path = "/"                             // Cookie berlaku untuk semua path

    // Mengirim cookie ke browser via header Set-Cookie
    http.SetCookie(w, cookie)
    fmt.Fprint(w, "Success Create Cookie")
}

// Handler untuk membaca cookie yang dikirim client
func GetCookie(w http.ResponseWriter, r *http.Request) {
    // r.Cookie() mencari cookie berdasarkan nama
    // Mengembalikan error jika cookie tidak ditemukan
    cookie, err := r.Cookie("X-Bani-Name")

    if err != nil {
        fmt.Fprint(w, "No Cookie")
    } else {
        name := cookie.Value
        fmt.Fprintf(w, "Hello %s", name)
    }
}

// Server dengan dua endpoint: set dan get cookie
func TestCookie(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/set-cookie", SetCookie)
    mux.HandleFunc("/get-cookie", GetCookie)

    server := http.Server{
        Addr:    "localhost:8080",
        Handler: mux,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

## Unit Test: Membuat Cookie

```go
func TestSetCookie(t *testing.T) {
    // Simulasi request dengan query parameter name=Bani
    req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Bani", nil)
    recorder := httptest.NewRecorder()

    SetCookie(recorder, req)

    // Ambil semua cookie yang ada di response
    cookies := recorder.Result().Cookies()

    for _, cookie := range cookies {
        fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
        // Output: Cookie X-Bani-Name: Bani
    }
}
```

### Cara Kerja TestSetCookie

```
httptest.NewRequest(GET, "/?name=Bani")
         ↓
SetCookie dipanggil → http.SetCookie() menulis header:
  Set-Cookie: X-Bani-Name=Bani; Path=/
         ↓
recorder.Result().Cookies() → membaca header Set-Cookie dari response
  → []*http.Cookie{ {Name: "X-Bani-Name", Value: "Bani"} }
```

## Unit Test: Membaca Cookie

```go
func TestGetCookie(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

    // Tambahkan cookie ke request simulasi secara manual
    // (meniru browser yang mengirim cookie secara otomatis)
    cookie := new(http.Cookie)
    cookie.Name = "X-Bani-Name"
    cookie.Value = "Burhanudin Rabbani"
    req.AddCookie(cookie)

    recorder := httptest.NewRecorder()
    GetCookie(recorder, req)

    body, _ := io.ReadAll(recorder.Result().Body)
    fmt.Println(string(body)) // Output: Hello Burhanudin Rabbani
}
```

### Cara Kerja TestGetCookie

```
req.AddCookie() menyisipkan header Cookie ke request:
  Cookie: X-Bani-Name=Burhanudin Rabbani
         ↓
GetCookie dipanggil → r.Cookie("X-Bani-Name") membaca header Cookie
  → &http.Cookie{Name: "X-Bani-Name", Value: "Burhanudin Rabbani"}
         ↓
fmt.Fprintf(w, "Hello %s", cookie.Value) → "Hello Burhanudin Rabbani"
```

## Alur Penggunaan

```
1. Client pertama kali request → server membuat cookie dengan http.SetCookie()
2. Browser menerima header Set-Cookie → menyimpan cookie
3. Request berikutnya → browser otomatis menyertakan cookie di header Cookie
4. Server membaca cookie → r.Cookie("nama") → memproses data sesi client
```

## Catatan Penting

| Catatan                                | Detail                                                                                |
| -------------------------------------- | ------------------------------------------------------------------------------------- |
| `r.Cookie()` bisa error                | Jika cookie tidak ditemukan, error `http.ErrNoCookie` dikembalikan — selalu cek error |
| `cookie.Path = "/"`                    | Menentukan path mana yang menerima cookie — `"/"` berarti semua endpoint              |
| Cookie bisa kedaluwarsa                | Gunakan `cookie.Expires` atau `cookie.MaxAge` untuk mengatur masa hidup cookie        |
| Cookie bukan tempat data sensitif      | Nilai cookie bisa dibaca user di browser — gunakan ID sesi, bukan data asli           |
| `req.AddCookie()` khusus untuk testing | Di browser nyata, cookie dikirim otomatis — `AddCookie` mensimulasikan perilaku itu   |

Next: [Fileserver](./14-fileserver.md)
