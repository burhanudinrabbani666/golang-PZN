# Template Caching

## Gambaran Umum

**Template caching** adalah teknik menyimpan template yang sudah di-parse ke dalam memori saat aplikasi pertama kali berjalan — sehingga setiap request tidak perlu melakukan parsing ulang dari file.

Tanpa caching, setiap kali handler dipanggil, Go harus membaca file dan mem-parse template dari awal. Ini sangat tidak efisien, terutama saat traffic tinggi.

---

## Tanpa Caching vs Dengan Caching

| Aspek               | Tanpa Caching                       | Dengan Caching               |
| ------------------- | ----------------------------------- | ---------------------------- |
| Parsing template    | Setiap request                      | Sekali saat startup          |
| Performa            | Lambat (I/O + parsing tiap request) | Cepat (langsung dari memori) |
| Penggunaan disk I/O | Setiap request                      | Hanya sekali                 |
| Cocok untuk         | Development/prototyping             | Production                   |

---

## Konsep: Embed + Global Variable

Kode ini menggabungkan dua teknik sekaligus:

```
//go:embed templates/*.gohtml
  → File template di-bundle langsung ke dalam binary Go (tidak perlu file di disk saat runtime)

var myTemplates = template.Must(...)
  → Template di-parse SEKALI saat program dimulai
  → Disimpan sebagai global variable (di memori)
  → Semua handler menggunakan instance yang sama
```

---

## Contoh Kode

```go
// go:embed meng-bundle semua file .gohtml ke dalam binary
// sehingga template tersedia meski file aslinya tidak ada di server
//go:embed templates/*.gohtml
var templates embed.FS

// myTemplates adalah global variable — template di-parse SEKALI saat startup
// template.Must() akan panic jika parsing gagal (lebih baik gagal saat startup
// daripada gagal diam-diam saat melayani request)
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

// TemplateCaching menggunakan template yang sudah di-cache
// Tidak ada parsing ulang di sini — langsung Execute dari memori
func TemplateCaching(w http.ResponseWriter, r *http.Request) {
    myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
    request  := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()

    TemplateCaching(recorder, request)

    body, _ := io.ReadAll(recorder.Result().Body)
    fmt.Println(string(body))
}
```

---

## Cara Kerja

```
Saat program PERTAMA KALI berjalan:

  //go:embed → baca semua file .gohtml → bundle ke binary
        ↓
  template.ParseFS() → parse semua template
        ↓
  template.Must()    → panic jika ada error (fail-fast saat startup)
        ↓
  myTemplates        → disimpan di memori sebagai global variable

Saat handler DIPANGGIL (setiap request):

  TemplateCaching dipanggil
        ↓
  myTemplates.ExecuteTemplate() → ambil template dari memori
        ↓
  Render dan kirim ke client
  (TIDAK ada I/O ke disk, TIDAK ada parsing)
```

---

## Mengapa `template.Must()`?

```go
// Tanpa Must — perlu cek error manual
tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
if err != nil {
    panic(err) // Harus ditangani sendiri
}

// Dengan Must — lebih ringkas, panic otomatis jika error
tmpl := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
```

`template.Must()` adalah helper yang menerima `(*Template, error)` — jika error tidak nil, ia langsung panic. Cocok untuk inisialisasi global karena error parsing template saat startup adalah kondisi yang tidak bisa di-recover.

---

## Alur Penggunaan Template Caching

```
1. Embed file template: //go:embed templates/*.gohtml
        ↓
2. Deklarasikan sebagai global variable di luar fungsi
        ↓
3. Parse sekali: template.Must(template.ParseFS(...))
        ↓
4. Di setiap handler, gunakan ExecuteTemplate() dari global variable
        ↓
5. Tidak perlu parsing lagi — template langsung diambil dari memori
```

---

## Catatan Penting

| Hal                         | Keterangan                                                                                              |
| --------------------------- | ------------------------------------------------------------------------------------------------------- |
| **Global variable = cache** | Dengan mendeklarasikan di luar fungsi, template hanya di-parse satu kali selama program berjalan        |
| **`template.Must()`**       | Panic saat startup jika template tidak valid — lebih baik gagal cepat daripada error saat melayani user |
| **`//go:embed`**            | Meng-bundle file ke dalam binary — file template tidak perlu ada di server secara terpisah              |
| **`embed.FS`**              | Filesystem virtual yang berisi file yang di-embed — digunakan bersama `template.ParseFS()`              |

---

Next: [XSS Cross Site Scripting](./22-xss-cross-site-scripting.md)
