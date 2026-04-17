# Params

## Gambaran Umum

| Aspek                        | Detail                                        |
| ---------------------------- | --------------------------------------------- |
| Topik                        | Path parameter di HttpRouter                  |
| Tipe data                    | `httprouter.Params`                           |
| Cara akses                   | `p.ByName("nama_param")`                      |
| Sintaks di route             | `:nama_param` di dalam URL path               |
| Perbedaan dengan query param | Path param ada di URL path, bukan setelah `?` |

`httprouter.Params` adalah parameter ketiga pada `httprouter.Handle`. Fungsinya adalah menampung nilai-nilai dinamis yang ada di dalam URL path. Berbeda dengan query parameter (`/products?id=1`), path parameter langsung tertanam di dalam struktur URL (`/products/1`) dan dideklarasikan saat mendaftarkan route.

## Path Parameter vs Query Parameter

| Aspek               | Path Parameter         | Query Parameter             |
| ------------------- | ---------------------- | --------------------------- |
| Contoh URL          | `/products/1`          | `/products?id=1`            |
| Dideklarasikan di   | Route: `/products/:id` | Tidak perlu deklarasi       |
| Cara akses          | `p.ByName("id")`       | `r.URL.Query().Get("id")`   |
| Dukungan ServeMux   | ✗ Tidak ada            | ✓ Ada                       |
| Dukungan HttpRouter | ✓ Ada                  | ✓ Ada                       |
| Cocok untuk         | ID resource, slug      | Filter, sorting, pagination |

## Contoh Kode

```go
func TestHttprouterParams(t *testing.T) {
	router := httprouter.New()

	// Tanda titik dua `:id` memberi tahu router bahwa segmen ini dinamis.
	// Nilai apapun yang ada di posisi tersebut akan ditangkap ke dalam Params
	// dengan kunci "id". Contoh: /products/1, /products/abc, /products/xyz-99
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// p.ByName("id") mengambil nilai dari segmen dinamis berdasarkan nama yang
		// didefinisikan di route. Jika URL-nya /products/1, maka hasilnya "1".
		text := "Products " + p.ByName("id")
		fmt.Fprint(w, text)
	})

	// Simulasikan request ke /products/1
	// Segmen "1" akan otomatis ditangkap sebagai nilai dari param "id"
	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Response body harus berisi "Products 1" karena p.ByName("id") = "1"
	assert.Equal(t, "Products 1", string(body))
}
```

## Cara Kerja

```
Pendaftaran Route
──────────────────────────────────────────────────────
router.GET("/products/:id", handler)
                    │
                    └── Router mencatat bahwa segmen ke-2
                        adalah dinamis dengan kunci "id"


Proses Request Masuk
──────────────────────────────────────────────────────
Request: GET /products/1
              │
              ▼
     Router mencocokkan pola /products/:id
              │
              ▼
     Params["id"] = "1"  ← nilai ditangkap otomatis
              │
              ▼
     handler(w, r, p) dipanggil
              │
              ▼
     p.ByName("id") → "1"
              │
              ▼
     Response: "Products 1"


Multiple Params
──────────────────────────────────────────────────────
Route:   /products/:category/:id
Request: /products/electronics/42

Params["category"] = "electronics"
Params["id"]       = "42"
```

## Alur Penggunaan

1. Saat mendaftarkan route, tandai segmen dinamis dengan awalan `:` — contoh: `/products/:id`
2. Router secara otomatis menangkap nilai segmen tersebut saat request masuk
3. Di dalam handler, akses nilai parameter lewat `p.ByName("nama_param")`
4. Gunakan nilai tersebut untuk query database, validasi, atau logika bisnis lainnya

## Catatan Penting

| Catatan                         | Penjelasan                                                                     |
| ------------------------------- | ------------------------------------------------------------------------------ |
| Nama param harus unik per route | Hindari dua segmen dinamis dengan nama sama di satu route                      |
| Nilai selalu bertipe `string`   | Lakukan konversi manual jika butuh `int` atau tipe lain (`strconv.Atoi`)       |
| Param tidak bisa kosong         | Jika segmen `:id` ada di route, URL tanpa nilai di posisi itu tidak akan cocok |
| Berbeda dengan wildcard `*`     | `:id` menangkap satu segmen; `*filepath` menangkap sisa seluruh path           |
| ServeMux tidak mendukung ini    | Fitur ini eksklusif di HttpRouter dan library routing lainnya                  |

Next: [Router Patterns](./04-router-patterns.md)
