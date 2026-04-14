# ServeMux

## Gambaran Umum

| Konsep             | Penjelasan                                                                |
| ------------------ | ------------------------------------------------------------------------- |
| `http.ServeMux`    | Router bawaan Go untuk memetakan URL ke handler                           |
| `NewServeMux()`    | Membuat instance ServeMux baru                                            |
| `HandleFunc()`     | Mendaftarkan fungsi handler untuk URL pattern tertentu                    |
| URL Pattern Exact  | Tanpa garis miring di akhir — hanya cocok tepat satu URL                  |
| URL Pattern Prefix | Dengan garis miring di akhir (`/images/`) — cocok semua URL berawalan itu |
| Prioritas Pattern  | Pattern yang lebih panjang selalu diprioritaskan                          |

`HandlerFunc` hanya bisa menangani satu endpoint. Jika kamu ingin menangani banyak URL sekaligus dalam satu server, gunakan **ServeMux** sebagai router utama.

## Penggunaan Dasar

```go
func TestServeMux(t *testing.T) {
    // Buat instance ServeMux baru
    mux := http.NewServeMux()

    // Daftarkan handler untuk root URL "/"
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello World")
    })

    // Daftarkan handler untuk URL "/hi"
    mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hi")
    })

    // Buat server dan pasang mux sebagai handler utama
    server := http.Server{
        Addr:    "localhost:8080",
        Handler: mux, // mux bertanggung jawab meneruskan request ke handler yang tepat
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

### Cara Kerja

```
Request masuk → http.Server → ServeMux → cocokkan URL Pattern → jalankan Handler
                                  |
                           "/" → "Hello World"
                          "/hi" → "Hi"
```

ServeMux bertindak sebagai **dispatcher**: ia membaca path dari setiap request yang masuk, lalu mencari handler yang terdaftar dengan pattern paling cocok.

## URL Pattern

Ada dua jenis URL pattern di ServeMux:

| Jenis              | Contoh     | Perilaku                                                      |
| ------------------ | ---------- | ------------------------------------------------------------- |
| Exact match        | `/hi`      | Hanya cocok untuk `/hi` saja                                  |
| Prefix match       | `/images/` | Cocok untuk `/images/`, `/images/foto`, `/images/a/b/c`, dst. |
| Fallback/catch-all | `/`        | Menangkap semua URL yang tidak cocok pattern lain             |

> **Aturan prioritas:** Jika dua pattern sama-sama cocok, ServeMux selalu memilih yang **lebih panjang (lebih spesifik)**.

```go
func TestServeMux(t *testing.T) {
    mux := http.NewServeMux()

    // Catch-all: menangani semua URL yang tidak cocok pattern lain
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello World")
    })

    // Exact match: hanya cocok untuk "/hi"
    mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hi")
    })

    // Prefix match: cocok untuk "/images/", "/images/kucing.png", dsb.
    mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Image")
    })

    // Prioritas lebih tinggi dari "/images/" karena lebih panjang
    // Cocok untuk "/images/thumbnails/", "/images/thumbnails/kecil.png", dsb.
    mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Thumbnails")
    })

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

### Cara Kerja Prioritas Pattern

```
Request: GET /images/thumbnails/kecil.png

ServeMux memeriksa semua pattern yang terdaftar:
  ✓ "/"                     → cocok (catch-all)
  ✓ "/images/"              → cocok (prefix match)
  ✓ "/images/thumbnails/"   → cocok (prefix match, lebih panjang)

Pemenang: "/images/thumbnails/" ← dipilih karena paling spesifik
```

```
Request: GET /images/foto.png

ServeMux memeriksa semua pattern:
  ✓ "/"          → cocok
  ✓ "/images/"   → cocok
  ✗ "/images/thumbnails/" → tidak cocok

Pemenang: "/images/" ← dipilih karena paling spesifik yang cocok
```

## Alur Penggunaan

```
1. Buat ServeMux        → http.NewServeMux()
2. Daftarkan handler    → mux.HandleFunc(pattern, handlerFunc)
3. Pasang ke server     → http.Server{Handler: mux}
4. Jalankan server      → server.ListenAndServe()
5. Request masuk        → mux mencocokkan URL → handler dieksekusi
```

## Catatan Penting

| Catatan                        | Detail                                                                                   |
| ------------------------------ | ---------------------------------------------------------------------------------------- |
| `/` adalah catch-all           | Semua URL yang tidak cocok pattern manapun akan diteruskan ke `/`                        |
| Garis miring di akhir          | Membuat pattern menjadi prefix — tanpa garis miring berarti exact match                  |
| Pattern lebih panjang menang   | `/images/thumbnails/` diprioritaskan di atas `/images/`                                  |
| Domain tidak perlu ditulis     | Pattern hanya berisi path, bukan full URL (`/images/` bukan `http://domain.com/images/`) |
| Tidak perlu urutan pendaftaran | ServeMux menentukan prioritas dari panjang pattern, bukan urutan `HandleFunc` dipanggil  |

Next: [Request](./7-request.md)
