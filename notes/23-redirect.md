# Redirect

## Gambaran Umum

| Aspek                 | Detail                     |
| --------------------- | -------------------------- |
| Topik                 | HTTP Redirect di Go        |
| Fungsi utama          | `http.Redirect()`          |
| Status code umum      | `301`, `302`, `307`, `308` |
| Header yang digunakan | `Location`                 |
| Paket                 | `net/http`                 |

HTTP Redirect adalah mekanisme standar yang meminta browser atau klien untuk berpindah ke URL lain. Dalam kehidupan nyata, redirect sering digunakan setelah proses login berhasil (arahkan ke dashboard), setelah submit form (hindari duplikasi data), atau saat URL lama diganti dengan yang baru.

Secara teknis, server cukup mengembalikan response dengan status code `3xx` dan menambahkan header `Location` yang berisi URL tujuan. Go menyederhanakan ini melalui fungsi `http.Redirect()`.

## Jenis Status Code Redirect

| Status Code | Nama               | Keterangan                                    |
| ----------- | ------------------ | --------------------------------------------- |
| `301`       | Moved Permanently  | URL lama tidak akan pernah digunakan lagi     |
| `302`       | Found              | Redirect sementara (method bisa berubah)      |
| `307`       | Temporary Redirect | Redirect sementara, method HTTP dipertahankan |
| `308`       | Permanent Redirect | Permanent, method HTTP dipertahankan          |

## Contoh Kode

```go
// Handler tujuan — halaman yang dituju setelah redirect
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	// Pengguna tiba di sini setelah diarahkan dari /redirect-from
	fmt.Fprint(writer, "Hello Redirect")
}

// Handler asal — memicu redirect ke URL internal lain
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// Setelah logika dijalankan (misal: validasi login),
	// arahkan pengguna ke /redirect-to sementara (307).
	// Browser akan tetap menggunakan method yang sama (GET/POST).
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

// Handler untuk redirect ke URL eksternal di luar aplikasi
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// Redirect ke website eksternal, misalnya profil GitHub.
	// Berguna untuk tautan afiliasi, SSO, atau referral.
	http.Redirect(writer, request, "https://github.com/burhanudinrabbani666", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()

	// Daftarkan ketiga route: asal, tujuan, dan eksternal
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-bani", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		panic(serverError)
	}
}
```

## Cara Kerja

```
Redirect Internal (/redirect-from → /redirect-to)
──────────────────────────────────────────────────

Browser          Server Go
  │                  │
  │  GET /redirect-from
  │─────────────────>│
  │                  │  Jalankan logika bisnis
  │                  │  http.Redirect(..., "/redirect-to", 307)
  │  307 + Location: /redirect-to
  │<─────────────────│
  │                  │
  │  GET /redirect-to  (otomatis oleh browser)
  │─────────────────>│
  │  200 "Hello Redirect"
  │<─────────────────│


Redirect Eksternal (/redirect-bani → github.com)
─────────────────────────────────────────────────

Browser          Server Go          GitHub
  │                  │                │
  │  GET /redirect-bani
  │─────────────────>│
  │  307 + Location: https://github.com/...
  │<─────────────────│
  │                  │
  │  GET https://github.com/...
  │────────────────────────────────>│
  │  200 (halaman GitHub)
  │<────────────────────────────────│
```

## Alur Penggunaan

1. Klien mengirim request ke URL asal (misal `/redirect-from`)
2. Handler menjalankan logika (validasi, autentikasi, dsb.)
3. `http.Redirect()` menulis header `Location` dan status `3xx` ke response
4. Browser/klien membaca status `3xx` dan header `Location`
5. Browser secara otomatis mengirim request baru ke URL tujuan
6. Handler tujuan memproses request dan mengirim response final

## Catatan Penting

| Catatan                                         | Penjelasan                                                                        |
| ----------------------------------------------- | --------------------------------------------------------------------------------- |
| Gunakan `307`/`308` untuk form POST             | `301`/`302` bisa mengubah method menjadi GET secara otomatis                      |
| Redirect hanya mengirim header                  | Tidak ada body yang diproses; logika berhenti setelah `http.Redirect()` dipanggil |
| Redirect eksternal bisa ke URL mana saja        | Tidak harus dalam domain yang sama                                                |
| Hindari redirect berantai                       | Terlalu banyak hop redirect memperlambat pengguna dan merusak SEO                 |
| `http.Redirect()` sudah memanggil `WriteHeader` | Jangan menulis header status lain setelah memanggilnya                            |

Next: [Upload File](./24-upload-file.md)
