# Server

## Gambaran Umum

**`http.Server`** adalah struct dari package `net/http` yang merepresentasikan sebuah web server di Go. Untuk membangun aplikasi web, langkah pertama adalah membuat dan menjalankan Server ini.

Server bertanggung jawab untuk:

- Mendengarkan koneksi masuk pada host dan port tertentu
- Meneruskan setiap request ke **Handler** yang sesuai
- Mengelola koneksi HTTP (keep-alive, timeout, dll.)

---

## Field Penting `http.Server`

| Field          | Tipe            | Keterangan                                                                       |
| -------------- | --------------- | -------------------------------------------------------------------------------- |
| `Addr`         | `string`        | Alamat server, format `"host:port"` — wajib diisi                                |
| `Handler`      | `http.Handler`  | Handler untuk memproses request — jika `nil`, menggunakan `http.DefaultServeMux` |
| `ReadTimeout`  | `time.Duration` | Batas waktu membaca request dari client                                          |
| `WriteTimeout` | `time.Duration` | Batas waktu menulis response ke client                                           |
| `IdleTimeout`  | `time.Duration` | Batas waktu koneksi idle (keep-alive)                                            |

---

## Contoh Kode

```go
func TestServer(t *testing.T) {
    server := http.Server{
        Addr: "localhost:8080", // Server berjalan di localhost, port 8080
    }

    // ListenAndServe() akan BLOCKING — program tidak akan lanjut ke baris berikutnya
    // sampai server dihentikan atau terjadi error
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

Setelah server berjalan, buka browser atau gunakan `curl` untuk mengaksesnya:

```bash
curl http://localhost:8080
```

---

## Cara Kerja

```
server.ListenAndServe()
        │
        ▼
Server mulai mendengarkan di localhost:8080
        │
        ▼
Client mengirim HTTP Request
        │
        ▼
Server menerima koneksi → teruskan ke Handler
        │
        ▼
Handler memproses request → kirim Response ke client
        │
        ▼
(Server terus berjalan, siap menerima request berikutnya)
```

---

## Contoh Server dengan Konfigurasi Lengkap

```go
server := http.Server{
    Addr:         "localhost:8080",
    ReadTimeout:  10 * time.Second, // Maksimal 10 detik untuk membaca request
    WriteTimeout: 10 * time.Second, // Maksimal 10 detik untuk menulis response
    IdleTimeout:  60 * time.Second, // Tutup koneksi idle setelah 60 detik
}
```

---

## Alur Penggunaan Server

```
1. Import package "net/http"
        ↓
2. Buat struct http.Server dengan minimal mengisi Addr
        ↓
3. (Opsional) Daftarkan Handler untuk memproses request
        ↓
4. Panggil server.ListenAndServe() — ini blocking, server mulai berjalan
        ↓
5. Tangani error jika server gagal berjalan (port sudah dipakai, dll.)
```

---

## Catatan Penting

| Hal                                      | Keterangan                                                                                                          |
| ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| **`ListenAndServe()` bersifat blocking** | Program akan berhenti di baris ini — server terus berjalan sampai dimatikan secara manual                           |
| **Tanpa Handler**                        | Jika `Handler` tidak diisi, server menggunakan `http.DefaultServeMux` sebagai default                               |
| **Port sudah dipakai**                   | Jika port 8080 sudah digunakan proses lain, `ListenAndServe()` akan mengembalikan error                             |
| **Typo di dokumentasi asli**             | `ListenAndServe()d` diperbaiki menjadi `ListenAndServe()`                                                           |
| **Konfigurasi timeout**                  | Selalu set `ReadTimeout` dan `WriteTimeout` untuk aplikasi production agar tidak rentan terhadap serangan slowloris |

---

Next: [Handler](./5-handler.md)
