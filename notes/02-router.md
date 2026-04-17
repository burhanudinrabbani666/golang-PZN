# Router

## Gambaran Umum

| Aspek                  | Detail                                            |
| ---------------------- | ------------------------------------------------- |
| Topik                  | Struct `Router` dari library HttpRouter           |
| Fungsi pembuatan       | `httprouter.New()`                                |
| Implementasi           | `http.Handler` (kompatibel dengan `http.Server`)  |
| Keunggulan vs ServeMux | Mendukung HTTP Method spesifik dan path parameter |
| Paket                  | `github.com/julienschmidt/httprouter`             |

`Router` adalah inti dari library HttpRouter. Karena mengimplementasikan interface `http.Handler`, router ini bisa langsung dipasang ke `http.Server` tanpa konfigurasi tambahan. Dibuat dengan `httprouter.New()` yang mengembalikan pointer `*Router`.

## Perbandingan: ServeMux vs HttpRouter

| Fitur                | ServeMux       | HttpRouter                   |
| -------------------- | -------------- | ---------------------------- |
| HTTP Method spesifik | ✗ Tidak bisa   | ✓ `GET`, `POST`, `PUT`, dll. |
| Path parameter       | ✗ Tidak ada    | ✓ `/user/:id`                |
| Handler type         | `http.Handler` | `httprouter.Handle`          |
| Performa routing     | Standar        | Lebih cepat (radix tree)     |

## httprouter.Handle vs http.Handler

ServeMux menggunakan `http.Handler` yang hanya memiliki dua parameter: `ResponseWriter` dan `*Request`. HttpRouter memperkenalkan tipe `httprouter.Handle` dengan **parameter ketiga**: `httprouter.Params`.

```go
// http.Handler — digunakan oleh ServeMux
type HandlerFunc func(http.ResponseWriter, *http.Request)

// httprouter.Handle — digunakan oleh HttpRouter
type Handle func(http.ResponseWriter, *http.Request, httprouter.Params)
//                                                    └── parameter tambahan
//                                                        untuk path parameter (:id, :name, dll.)
```

Parameter `Params` memungkinkan kita mengambil nilai dari segmen URL dinamis seperti `/user/:id`. Ini akan dibahas lebih lanjut di chapter Params.

## Contoh Kode

```go
func TestHttprouter(t *testing.T) {
	// Buat router baru — ini adalah pengganti ServeMux
	router := httprouter.New()

	// Daftarkan route GET "/" dengan handler bertipe httprouter.Handle.
	// Perhatikan parameter ketiga `p httprouter.Params` — ini yang membedakan
	// dari http.HandlerFunc biasa.
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	// Simulasikan request GET ke "/" menggunakan httptest
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Router mengimplementasikan http.Handler, jadi bisa dipanggil dengan ServeHTTP
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Verifikasi bahwa response body sesuai ekspektasi
	assert.Equal(t, "Hello World", string(body))
}
```

## Cara Kerja

```
Inisialisasi Router
────────────────────────────────────────────
httprouter.New()
       │
       ▼
  *Router (radix tree kosong)
       │
       │  router.GET("/", handler)
       │  router.POST("/login", handler)
       ▼
  Routing table terisi


Alur Request Masuk
────────────────────────────────────────────
Request GET /
       │
       ▼
  router.ServeHTTP()
       │
       ▼
  Cocokkan Method + Path di radix tree
       │
       ├── Cocok → jalankan handler(w, r, params)
       │
       └── Tidak cocok → 404 / 405 Method Not Allowed
```

## Alur Penggunaan

1. Import package `httprouter` ke dalam project
2. Buat router dengan `httprouter.New()`
3. Daftarkan route menggunakan method yang sesuai: `router.GET()`, `router.POST()`, dll.
4. Pasang router ke `http.Server` melalui field `Handler`
5. Jalankan server — router siap menerima request

## Catatan Penting

| Catatan                                  | Penjelasan                                                                                      |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `httprouter.New()` mengembalikan pointer | Gunakan `*Router`, bukan value                                                                  |
| Method tidak cocok → `405`               | Jika route ada tapi method salah, router otomatis balas `405 Method Not Allowed`                |
| Parameter `Params` bisa diabaikan        | Jika route tidak memiliki path parameter, `p` tidak digunakan — tapi tetap harus dideklarasikan |
| Kompatibel penuh dengan `http.Server`    | Cukup set `Handler: router` di struct `http.Server`                                             |

Next: [Params](./03-params.md)
