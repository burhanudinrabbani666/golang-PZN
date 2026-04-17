# Serve File

## Gambaran Umum

| Aspek                 | Detail                                              |
| --------------------- | --------------------------------------------------- |
| Topik                 | Melayani file statis dengan HttpRouter              |
| Fungsi utama          | `router.ServeFiles(path, filesystem)`               |
| Syarat parameter path | Harus menggunakan catch-all parameter (`*filepath`) |
| Sumber file           | Folder di disk atau `embed.FS` (Go Embed)           |
| Paket pendukung       | `io/fs`, `embed`, `net/http`                        |

HttpRouter menyediakan fungsi `ServeFiles` untuk melayani file statis seperti HTML, CSS, gambar, atau file teks. Fungsi ini membutuhkan catch-all parameter di path karena nama file bisa bervariasi dan bisa berada di subfolder manapun.

## Dua Cara Menyediakan FileSystem

| Cara           | Implementasi                  | Keterangan                                  |
| -------------- | ----------------------------- | ------------------------------------------- |
| Folder di disk | `http.Dir("./static")`        | File dibaca langsung dari disk saat runtime |
| Go Embed       | `http.FS(dir)` + `//go:embed` | File dikompilasi ke dalam binary saat build |

Go Embed lebih direkomendasikan untuk production karena file ikut terbundle di dalam binary — tidak perlu mendistribusikan folder terpisah.

## Contoh Kode

```go
// Direktif embed memberitahu compiler untuk menyertakan folder "resources"
// beserta seluruh isinya ke dalam binary saat proses build.
// Variabel `resources` bertipe embed.FS dan bersifat read-only.
//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	// fs.Sub membuat "sub-filesystem" yang root-nya dimulai dari folder "resources".
	// Ini diperlukan agar URL /files/hello.txt langsung memetakan ke resources/hello.txt,
	// bukan resources/resources/hello.txt.
	directory, _ := fs.Sub(resources, "resources")

	// ServeFiles mendaftarkan handler untuk melayani file statis.
	// Parameter pertama HARUS menggunakan catch-all (*filepath) karena
	// nama dan path file bisa berbeda-beda di setiap request.
	// Parameter kedua adalah FileSystem yang akan dibaca.
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// Simulasikan request untuk mengambil file hello.txt
	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Isi file resources/hello.txt adalah "Hello httprouter"
	assert.Equal(t, "Hello httprouter", string(body))
}
```

## Cara Kerja

```
Struktur Folder Proyek
──────────────────────────────────────────────────────
project/
└── resources/          ← di-embed ke dalam binary
    ├── hello.txt
    └── images/
        └── logo.png


Alur Request File
──────────────────────────────────────────────────────
Request: GET /files/hello.txt
              │
              ▼
    router.ServeFiles menangkap *filepath = "/hello.txt"
              │
              ▼
    Cari file di embed.FS → resources/hello.txt
              │
              ├── Ditemukan → kirim isi file sebagai response body
              │
              └── Tidak ditemukan → 404 Not Found


Pemetaan URL ke File
──────────────────────────────────────────────────────
/files/*filepath        →   embed.FS (root: "resources")
/files/hello.txt        →   resources/hello.txt
/files/images/logo.png  →   resources/images/logo.png
```

## Alur Penggunaan

1. Buat folder untuk menyimpan file statis (misal `resources/`)
2. Tambahkan direktif `//go:embed resources` di atas variabel `embed.FS`
3. Gunakan `fs.Sub()` untuk menentukan root folder agar path URL tidak redundan
4. Daftarkan route dengan `router.ServeFiles("/prefix/*filepath", http.FS(dir))`
5. Akses file lewat URL: `/prefix/nama-file.ext`

## Catatan Penting

| Catatan                                           | Penjelasan                                                              |
| ------------------------------------------------- | ----------------------------------------------------------------------- |
| `*filepath` wajib ada di path                     | `ServeFiles` akan panic jika path tidak mengandung catch-all parameter  |
| `fs.Sub()` penting untuk path yang benar          | Tanpa `fs.Sub`, root FS dimulai dari root embed — path bisa tidak cocok |
| File embed bersifat read-only                     | Tidak bisa menulis atau menghapus file di `embed.FS` saat runtime       |
| Gunakan `http.Dir("./folder")` untuk file dinamis | Jika file perlu diupdate tanpa rebuild, gunakan folder di disk          |
| Direktif embed harus tepat sebelum variabel       | `//go:embed resources` dan `var resources embed.FS` harus berurutan     |

Next: [Panic Handler](./06-panic-handler.md)
