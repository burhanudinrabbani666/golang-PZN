# ServeFile

## Gambaran Umum

| Konsep                     | Penjelasan                                                                 |
| -------------------------- | -------------------------------------------------------------------------- |
| `http.ServeFile()`         | Mengirim satu file tertentu sebagai HTTP response                          |
| `http.FileServer()`        | Melayani seluruh direktori secara otomatis berdasarkan URL                 |
| `//go:embed` pada `string` | Menyematkan isi file teks langsung ke dalam variabel string saat kompilasi |
| `fmt.Fprint(w, isi)`       | Menulis isi file yang sudah di-embed langsung ke response                  |

`http.ServeFile()` digunakan saat kita ingin **menentukan sendiri file mana yang dikirim** sebagai response, berdasarkan logika tertentu — bukan memetakan URL ke file secara otomatis seperti `FileServer`.

## Melayani File Berdasarkan Kondisi

```go
func ServeFile(w http.ResponseWriter, r *http.Request) {
    // Cek apakah query parameter "name" ada
    if r.URL.Query().Get("name") != "" {
        // Kirim file ok.html jika parameter tersedia
        http.ServeFile(w, r, "./resources/ok.html")
    } else {
        // Kirim file notFound.html jika parameter tidak ada
        http.ServeFile(w, r, "./resources/notFound.html")
    }
}

func TestServeFileServer(t *testing.T) {
    server := http.Server{
        Addr:    "localhost:8080",
        Handler: http.HandlerFunc(ServeFile),
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

### Cara Kerja

```
Request: GET /?name=Bani
  r.URL.Query().Get("name") → "Bani" (tidak kosong)
  http.ServeFile(w, r, "./resources/ok.html")
  → Go membuka file, membaca isinya, menulis ke response
  → Browser menerima konten ok.html

Request: GET /
  r.URL.Query().Get("name") → "" (kosong)
  http.ServeFile(w, r, "./resources/notFound.html")
  → Browser menerima konten notFound.html
```

## ServeFile dengan Go Embed

`http.ServeFile()` hanya menerima path string ke filesystem, sehingga tidak bisa langsung diintegrasikan dengan `embed.FS`. Solusinya adalah menyematkan file sebagai variabel `string` menggunakan `//go:embed`, lalu menulisnya langsung ke response menggunakan `fmt.Fprint()`.

```go
// Sematkan isi file ok.html ke dalam variabel string saat kompilasi
//go:embed resources/ok.html
var resourcesOK string

// Sematkan isi file notFound.html ke dalam variabel string saat kompilasi
//go:embed resources/notFound.html
var resourcesNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
    if r.URL.Query().Get("name") != "" {
        // Tulis langsung isi file yang sudah tersimpan di variabel
        fmt.Fprint(w, resourcesOK)
    } else {
        fmt.Fprint(w, resourcesNotFound)
    }
}

func TestServeFileServerEmbed(t *testing.T) {
    server := http.Server{
        Addr:    "localhost:8080",
        Handler: http.HandlerFunc(ServeFileEmbed),
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

### Cara Kerja

```
Saat kompilasi (go build):
  //go:embed resources/ok.html
  → isi file dibaca dan disimpan sebagai string di dalam binary
  → tidak perlu file fisik saat runtime

Saat request masuk:
  fmt.Fprint(w, resourcesOK)
  → isi HTML langsung ditulis ke response dari memori
  → tidak ada pembacaan file dari disk
```

## Perbandingan: ServeFile vs ServeFile Embed

|                              | `http.ServeFile()`                   | Go Embed + `fmt.Fprint()`                   |
| ---------------------------- | ------------------------------------ | ------------------------------------------- |
| Sumber file                  | Dibaca dari disk saat request        | Sudah ada di memori sejak binary dijalankan |
| Perlu file di server?        | Ya                                   | Tidak                                       |
| Bisa pakai logika kondisi?   | Ya                                   | Ya                                          |
| Integrasi dengan `embed.FS`? | Tidak langsung                       | Ya, via variabel `string`                   |
| Cocok untuk                  | File yang bisa berubah tanpa rebuild | Distribusi single binary                    |

## Perbandingan: ServeFile vs FileServer

|                      | `http.ServeFile()`                       | `http.FileServer()`               |
| -------------------- | ---------------------------------------- | --------------------------------- |
| Kontrol file         | Manual — kita tentukan sendiri file mana | Otomatis — berdasarkan URL path   |
| Cocok untuk          | Kondisi dinamis, routing file kustom     | Melayani seluruh direktori statis |
| Perlu `StripPrefix`? | Tidak                                    | Ya, jika ada prefix URL           |

## Alur Penggunaan

```
ServeFile biasa:
  1. Terima request
  2. Tentukan file mana yang akan dikirim berdasarkan logika
  3. Panggil http.ServeFile(w, r, "path/ke/file")
  4. Go membuka file dan mengirimnya sebagai response

ServeFile dengan Embed:
  1. Sematkan file dengan //go:embed ke variabel string
  2. Terima request
  3. Tentukan variabel mana yang akan dikirim
  4. Panggil fmt.Fprint(w, variabel)
```

## Catatan Penting

| Catatan                                        | Detail                                                                                                                  |
| ---------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `http.ServeFile()` otomatis set `Content-Type` | Go mendeteksi tipe file dari ekstensi dan mengisi header `Content-Type` secara otomatis                                 |
| `fmt.Fprint()` tidak set `Content-Type`        | Saat menggunakan embed + `fmt.Fprint`, set header manual jika diperlukan: `w.Header().Set("Content-Type", "text/html")` |
| Embed `string` hanya untuk file teks           | Untuk file biner (gambar, font), gunakan `[]byte` bukan `string`                                                        |
| `//go:embed` path relatif terhadap file `.go`  | Path embed ditulis relatif dari lokasi file Go yang mendefinisikan variabel tersebut                                    |

Next: [Template](./16-template.md)
