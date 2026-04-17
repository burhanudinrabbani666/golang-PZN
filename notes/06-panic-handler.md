# Panic Handler

## Gambaran Umum

| Aspek                    | Detail                                          |
| ------------------------ | ----------------------------------------------- |
| Topik                    | Menangani panic di HttpRouter                   |
| Atribut                  | `router.PanicHandler`                           |
| Tipe fungsi              | `func(http.ResponseWriter, *http.Request, any)` |
| Parameter ketiga (`any`) | Nilai yang dikirim saat `panic()` dipanggil     |
| Tanpa PanicHandler       | Server crash, response tidak dikirim ke client  |

Ketika sebuah handler memanggil `panic()`, secara default goroutine tersebut akan berhenti dan server tidak mengirimkan response apapun ke client. HttpRouter menyediakan atribut `PanicHandler` sebagai jaring pengaman — mirip seperti `recover()` yang dibungkus middleware, tetapi sudah tersedia built-in tanpa perlu kode tambahan.

## Tanpa vs Dengan PanicHandler

| Kondisi               | Perilaku Server                   | Response ke Client                 |
| --------------------- | --------------------------------- | ---------------------------------- |
| Tanpa `PanicHandler`  | Goroutine crash, log ke stderr    | Koneksi terputus / `500` kosong    |
| Dengan `PanicHandler` | Panic dicegat, handler dijalankan | Response custom sesuai logika kita |

## Contoh Kode

```go
func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// PanicHandler dipanggil otomatis oleh router setiap kali ada panic
	// di handler manapun. Parameter `i` berisi nilai yang dikirim ke panic(),
	// bisa berupa string, error, atau tipe apapun.
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i any) {
		// Di production, gunakan bagian ini untuk:
		// - Mencatat error ke logging system (misal: Sentry, Datadog)
		// - Mengirim response 500 dengan pesan yang aman untuk user
		// - Memastikan client tidak mendapat response kosong
		fmt.Fprint(w, "Panic: ", i)
	}

	// Handler ini sengaja memanggil panic untuk mensimulasikan error tak terduga,
	// misalnya nil pointer dereference atau assertion yang gagal.
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("ups") // nilai "ups" akan diterima sebagai parameter `i` di PanicHandler
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// PanicHandler menangkap panic dan menulis "Panic: ups" ke response
	assert.Equal(t, "Panic: ups", string(body))
}
```

## Cara Kerja

```
Alur Normal (tanpa panic)
──────────────────────────────────────────────────────
Request → router.ServeHTTP → handler(w, r, p) → Response


Alur dengan Panic
──────────────────────────────────────────────────────
Request → router.ServeHTTP → handler(w, r, p)
                                    │
                                  panic("ups")
                                    │
                                    ▼
                          Router mencegat panic
                          (via recover() internal)
                                    │
                                    ▼
                          PanicHandler(w, r, "ups")
                                    │
                                    ▼
                          Response dikirim ke client
                          (server tetap berjalan)


Tanpa PanicHandler
──────────────────────────────────────────────────────
Request → handler → panic("ups")
                         │
                         ▼
                  Goroutine crash
                  Log ke stderr
                  Client: koneksi terputus
                  Server: tetap jalan tapi request ini gagal
```

## Alur Penggunaan

1. Buat router dengan `httprouter.New()`
2. Tetapkan `router.PanicHandler` dengan fungsi penanganan kustom
3. Di dalam `PanicHandler`, gunakan parameter `i` untuk mendapatkan nilai panic
4. Tulis response yang sesuai — umumnya status `500` dengan pesan error
5. Opsional: kirim log ke sistem monitoring sebelum membalas client

## Catatan Penting

| Catatan                                            | Penjelasan                                                                    |
| -------------------------------------------------- | ----------------------------------------------------------------------------- |
| `PanicHandler` mencegah server crash total         | Goroutine yang panic dihentikan, tapi server tetap melayani request lain      |
| Parameter `i` bertipe `any`                        | Lakukan type assertion jika butuh mengakses nilai spesifik, misal `i.(error)` |
| Satu `PanicHandler` berlaku untuk semua route      | Tidak perlu mendaftarkan per-route                                            |
| Tetapkan sebelum mendaftarkan route                | Meski urutan tidak berpengaruh secara teknis, ini lebih mudah dibaca          |
| Di production, jangan kirim detail panic ke client | Log secara internal; kirim pesan generik ke user untuk alasan keamanan        |

Next: [Not Found Handler](./07-not-found-handler.md)
