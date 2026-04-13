# Header

## Gambaran Umum

**Header** adalah informasi tambahan yang dikirim bersama HTTP Request maupun Response. Header berisi metadata seperti tipe konten, autentikasi, informasi browser, dan lain-lain — bukan bagian dari URL atau body.

```
HTTP Request:
  GET /api/users HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json    ← ini Header
  Authorization: Bearer token123    ← ini Header

HTTP Response:
  HTTP/1.1 200 OK
  Content-Type: application/json   ← ini Header
  X-Powered-By: Burhanudin Rabbani ← ini Header (custom)
```

---

## Request Header vs Response Header

| Aspek     | Request Header                     | Response Header                |
| --------- | ---------------------------------- | ------------------------------ |
| Arah      | Client → Server                    | Server → Client                |
| Cara baca | `r.Header.Get("key")`              | `response.Header.Get("key")`   |
| Cara set  | `request.Header.Add("key", "val")` | `w.Header().Add("key", "val")` |
| Contoh    | `Content-Type`, `Authorization`    | `Content-Type`, `X-Powered-By` |

---

## Header vs Query Parameter

| Aspek              | Header                                   | Query Parameter               |
| ------------------ | ---------------------------------------- | ----------------------------- |
| Lokasi             | Di metadata HTTP (tidak terlihat di URL) | Di URL (`?key=value`)         |
| Case sensitive key | ❌ Tidak (case-insensitive)              | ✅ Ya (case-sensitive)        |
| Cocok untuk        | Autentikasi, metadata, token             | Filter, pencarian, pagination |

---

## 1. Request Header — Membaca dari Client

Header request dibaca melalui `r.Header.Get("key")`. Key bersifat **case-insensitive** — `"content-type"` dan `"Content-Type"` menghasilkan nilai yang sama.

```go
// RequestHeader mengambil nilai Content-Type dari header request
// dan mengirimkannya kembali sebagai body response
func RequestHeader(w http.ResponseWriter, r *http.Request) {
    // Header.Get() otomatis menormalisasi key menjadi canonical form
    // "content-type", "CONTENT-TYPE", "Content-Type" → semua sama
    contentType := r.Header.Get("Content-Type")
    fmt.Fprint(w, contentType)
}

func TestReqHeader(t *testing.T) {
    request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
    request.Header.Add("Content-Type", "application/json") // Set header di request

    recorder := httptest.NewRecorder()
    RequestHeader(recorder, request)

    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)

    fmt.Println(string(body)) // Output: application/json
}
```

---

## 2. Response Header — Mengirim ke Client

Header response ditambahkan melalui `w.Header().Add("key", "value")`. Header **harus ditambahkan sebelum** menulis body response.

```go
// ResponseHeader menambahkan custom header ke response
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
    // Tambahkan header SEBELUM menulis body
    // Header yang ditambahkan setelah Fprint/Write tidak akan terkirim
    w.Header().Add("X-Powered-By", "Burhanudin Rabbani")
    fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
    request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)

    recorder := httptest.NewRecorder()
    ResponseHeader(recorder, request)

    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)

    fmt.Println(string(body))                          // Output: OK
    fmt.Println(response.Header.Get("x-powered-by"))  // Output: Burhanudin Rabbani
    // Perhatikan: key "x-powered-by" (lowercase) tetap bisa membaca
    // header "X-Powered-By" karena header bersifat case-insensitive
}
```

---

## Cara Kerja

```
Request Header:

  Client mengirim request dengan header:
    Content-Type: application/json
          │
  r.Header → map[string][]string
    {"Content-Type": ["application/json"]}
          │
  r.Header.Get("Content-Type") → "application/json"

Response Header:

  w.Header().Add("X-Powered-By", "Burhanudin Rabbani")
          │
  Header ditambahkan ke response sebelum dikirim
          │
  Client menerima response:
    HTTP/1.1 200 OK
    X-Powered-By: Burhanudin Rabbani
```

---

## Method Header yang Tersedia

| Method                   | Fungsi                                                                 |
| ------------------------ | ---------------------------------------------------------------------- |
| `Header.Get(key)`        | Mengambil nilai pertama dari header dengan key tersebut                |
| `Header.Add(key, value)` | Menambahkan nilai header — bisa ada banyak nilai untuk satu key        |
| `Header.Set(key, value)` | Menetapkan nilai header — menggantikan nilai sebelumnya jika sudah ada |
| `Header.Del(key)`        | Menghapus header dengan key tersebut                                   |

---

## Alur Penggunaan Header

```
Membaca Request Header:
  1. Client mengirim request dengan header
  2. Akses di handler: r.Header.Get("nama-header")
  3. Gunakan nilainya untuk logika bisnis

Menulis Response Header:
  1. Di dalam handler, panggil w.Header().Add("key", "value")
  2. Pastikan SEBELUM menulis body (fmt.Fprint, w.Write, dll.)
  3. Client menerima response beserta header yang ditambahkan
```

---

## Catatan Penting

| Hal                     | Keterangan                                                                                                   |
| ----------------------- | ------------------------------------------------------------------------------------------------------------ |
| **Case-insensitive**    | Key header tidak case-sensitive — `"content-type"` sama dengan `"Content-Type"`                              |
| **Header sebelum body** | Response header **wajib** ditambahkan sebelum menulis body — setelah itu tidak bisa diubah                   |
| **`Add` vs `Set`**      | `Add` menambahkan nilai baru (bisa banyak nilai untuk satu key); `Set` menggantikan semua nilai lama         |
| **Prefix `X-`**         | Header custom biasanya diawali `X-` (seperti `X-Powered-By`) — konvensi untuk membedakan dari header standar |
| **Koreksi typo**        | `ResponesHeader` diperbaiki menjadi `ResponseHeader`                                                         |

---

Next: [Form Post](./11-form-post.md)
