# Handler

## Gambaran Umum

**Handler** adalah komponen yang bertugas menerima dan memproses HTTP request yang masuk ke server. Jika `http.Server` adalah "pintu gerbang" yang menerima koneksi, maka Handler adalah "penerima tamu" yang memutuskan apa yang harus dilakukan dengan setiap request.

Di Go, Handler direpresentasikan sebagai **interface**:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Setiap tipe yang mengimplementasikan method `ServeHTTP()` bisa digunakan sebagai Handler.

---

## Server vs Handler

| Komponen       | Tanggung Jawab                                       |
| -------------- | ---------------------------------------------------- |
| `http.Server`  | Mendengarkan koneksi masuk di host dan port tertentu |
| `http.Handler` | Memproses setiap HTTP request dan menulis response   |

---

## `http.HandlerFunc`

**`HandlerFunc`** adalah cara paling praktis untuk membuat Handler. Ia adalah tipe fungsi yang secara otomatis mengimplementasikan interface `Handler` — sehingga kita bisa langsung menggunakan fungsi biasa sebagai Handler tanpa perlu membuat struct baru.

```go
// Signature yang harus dipenuhi oleh HandlerFunc
type HandlerFunc func(ResponseWriter, *Request)

// HandlerFunc mengimplementasikan interface Handler melalui method ServeHTTP
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

---

## Contoh Kode

```go
func TestHandler(t *testing.T) {
    // Deklarasi HandlerFunc — fungsi ini dipanggil setiap ada request masuk
    var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
        // w (ResponseWriter) digunakan untuk menulis response ke client
        // r (*Request) berisi semua informasi tentang request yang masuk
        fmt.Fprint(w, "Hello World")
    }

    server := http.Server{
        Addr:    "localhost:8080",
        Handler: handler, // Daftarkan handler ke server
    }

    // ListenAndServe bersifat blocking — server terus berjalan
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

---

## Cara Kerja

```
Client mengirim HTTP Request
        ↓
http.Server menerima koneksi di localhost:8080
        ↓
Server memanggil handler.ServeHTTP(w, r)
        ↓
HandlerFunc dieksekusi:
  - Baca data dari r (*http.Request) jika perlu
  - Tulis response ke w (http.ResponseWriter)
  - fmt.Fprint(w, "Hello World")
        ↓
Response dikirim kembali ke client
```

---

## Parameter Handler

| Parameter | Tipe                  | Fungsi                                                       |
| --------- | --------------------- | ------------------------------------------------------------ |
| `w`       | `http.ResponseWriter` | Menulis response — status code, header, dan body             |
| `r`       | `*http.Request`       | Membaca request — URL, method, header, body, query parameter |

---

## Alur Penggunaan Handler

```
1. Buat fungsi dengan signature: func(http.ResponseWriter, *http.Request)
        ↓
2. Cast ke http.HandlerFunc atau langsung gunakan sebagai handler
        ↓
3. Daftarkan ke http.Server melalui field Handler
        ↓
4. Jalankan server dengan ListenAndServe()
        ↓
5. Setiap request yang masuk akan memanggil fungsi handler tersebut
```

---

## Catatan Penting

| Hal                                   | Keterangan                                                                                                 |
| ------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| **Interface `Handler`**               | Apapun yang punya method `ServeHTTP(ResponseWriter, *Request)` bisa jadi Handler                           |
| **`HandlerFunc` adalah shortcut**     | Mengubah fungsi biasa menjadi Handler tanpa perlu membuat struct                                           |
| **Satu handler untuk semua path**     | Contoh di atas menangani semua URL — gunakan `ServeMux` untuk routing ke path yang berbeda                 |
| **`fmt.Fprint(w, ...)` menulis body** | `ResponseWriter` mengimplementasikan `io.Writer` — bisa digunakan dengan semua fungsi yang menerima Writer |

---

Next: [ServeMux](./6-servemux.md)
