# File Server

## Gambaran Umum

| Konsep               | Penjelasan                                                              |
| -------------------- | ----------------------------------------------------------------------- |
| `http.FileServer()`  | Membuat Handler yang melayani file statis dari sebuah direktori         |
| `http.Dir("path")`   | Menunjuk direktori di filesystem sebagai sumber file                    |
| `http.StripPrefix()` | Menghapus prefix URL sebelum diteruskan ke FileServer                   |
| `embed.FS`           | Menyematkan file ke dalam binary saat kompilasi (Go 1.16+)              |
| `fs.Sub()`           | Mengambil sub-direktori dari `embed.FS` agar path lebih bersih          |
| `http.FS()`          | Mengubah `fs.FS` menjadi `http.FileSystem` yang bisa dipakai FileServer |

`http.FileServer` adalah Handler siap pakai yang melayani file statis (HTML, CSS, JS, gambar, dsb.) langsung dari direktori. Kita tidak perlu membuka dan mengirim file secara manual — cukup daftarkan handler-nya ke server atau ServeMux.

## FileServer dari Direktori Lokal

```go
func TestFileServer(t *testing.T) {
    // Tentukan direktori sumber file statis
    directory := http.Dir("./resources")

    // Buat FileServer Handler dari direktori tersebut
    fileServer := http.FileServer(directory)

    mux := http.NewServeMux()

    // StripPrefix wajib digunakan agar prefix "/static" tidak ikut dicari sebagai folder
    // Tanpa ini, request ke /static/index.js akan mencari file di ./resources/static/index.js
    // Dengan StripPrefix, prefix "/static" dihapus dulu → FileServer mencari ./resources/index.js
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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

### Cara Kerja StripPrefix

```
Struktur folder:
  resources/
    index.js
    style.css

Tanpa StripPrefix:
  Request: GET /static/index.js
  FileServer mencari: ./resources/static/index.js  ← tidak ada → 404 Not Found

Dengan StripPrefix("/static", fileServer):
  Request: GET /static/index.js
  StripPrefix menghapus "/static" → /index.js
  FileServer mencari: ./resources/index.js          ← ada → 200 OK
```

## FileServer dengan Go Embed

Go 1.16 memperkenalkan fitur **embed** yang memungkinkan file statis disematkan langsung ke dalam binary saat kompilasi. Artinya file tidak perlu ikut didistribusikan secara terpisah — semuanya sudah ada di dalam satu file executable.

```go
// Direktif embed: file dalam folder "resources" akan disematkan ke binary
//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
    // fs.Sub() mengambil sub-direktori "resources" agar path URL tidak menyertakan nama folder
    // Tanpa ini, file index.js harus diakses via /resources/index.js
    // Dengan fs.Sub(), file langsung bisa diakses via /index.js
    directory, _ := fs.Sub(resources, "resources")

    // http.FS() mengubah embed.FS menjadi http.FileSystem yang dikenali FileServer
    fileServer := http.FileServer(http.FS(directory))

    mux := http.NewServeMux()
    mux.Handle("/", fileServer)

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

### Cara Kerja fs.Sub

```
Struktur folder embed:
  resources/
    index.js
    style.css

Tanpa fs.Sub():
  embed.FS menyimpan: "resources/index.js"
  Request: GET /index.js  → 404 Not Found
  Request: GET /resources/index.js  → 200 OK  ← perlu path panjang

Dengan fs.Sub(resources, "resources"):
  Sub-direktori menjadi root baru
  Request: GET /index.js  → 200 OK  ← langsung akses tanpa prefix folder
```

## Perbandingan: FileServer Biasa vs Go Embed

|                        | FileServer Biasa             | Go Embed                                     |
| ---------------------- | ---------------------------- | -------------------------------------------- |
| Sumber file            | Direktori di filesystem      | Disematkan dalam binary                      |
| Perlu distribusi file? | Ya, file harus ada di server | Tidak, sudah ada di dalam binary             |
| Gunakan `StripPrefix`? | Ya, jika ada prefix URL      | Tidak diperlukan                             |
| Gunakan `fs.Sub()`?    | Tidak                        | Ya, untuk menghilangkan nama folder dari URL |
| Cocok untuk            | Development, file besar      | Distribusi aplikasi single binary            |

## Alur Penggunaan

```
FileServer Biasa:
  1. Siapkan folder berisi file statis (./resources)
  2. Buat handler → http.FileServer(http.Dir("./resources"))
  3. Daftarkan ke mux dengan StripPrefix jika URL punya prefix
  4. Server melayani file sesuai URL yang diminta

Go Embed:
  1. Tambahkan direktif //go:embed resources di atas variabel embed.FS
  2. Ambil sub-direktori → fs.Sub(resources, "resources")
  3. Buat handler → http.FileServer(http.FS(directory))
  4. Daftarkan ke mux — file langsung bisa diakses sesuai nama file
```

## Catatan Penting

| Catatan                                       | Detail                                                                                           |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------ |
| `StripPrefix` wajib jika ada prefix URL       | Tanpanya, FileServer akan mencari folder yang namanya sama dengan prefix URL                     |
| `//go:embed` harus tepat di atas variabel     | Direktif embed tidak boleh ada baris kosong antara komentar dan deklarasi variabel               |
| `fs.Sub()` mengembalikan error                | Selalu tangani error dari `fs.Sub()` meskipun dalam testing                                      |
| FileServer melayani directory listing         | Jika tidak ada `index.html`, FileServer akan menampilkan daftar file — pertimbangkan keamanannya |
| Go Embed tidak bisa embed file di luar module | Path embed harus berada di dalam direktori yang sama atau sub-direktori modul Go                 |

Next: [Servefile](./15-servefile.md)
