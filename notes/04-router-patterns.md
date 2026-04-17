# Router Patterns

## Gambaran Umum

| Aspek           | Detail                                                             |
| --------------- | ------------------------------------------------------------------ |
| Topik           | Pola parameter URL di HttpRouter                                   |
| Jenis pola      | Named Parameter (`:nama`) dan Catch-All Parameter (`*nama`)        |
| Perbedaan utama | Named menangkap satu segmen; Catch-All menangkap sisa seluruh path |
| Cara akses      | `p.ByName("nama")` untuk keduanya                                  |

HttpRouter mendukung dua pola pembuatan parameter URL. Memilih pola yang tepat penting karena menentukan URL mana yang akan cocok (match) dan mana yang tidak.

## Named Parameter

Named parameter menangkap **tepat satu segmen** URL. Ditulis dengan awalan `:` diikuti nama parameter. Satu route bisa memiliki beberapa named parameter selama berada di segmen yang berbeda.

| Pattern       | URL                 | Hasil                      |
| ------------- | ------------------- | -------------------------- |
| `/user/:user` | `/user/eko`         | ✓ match, `user = "eko"`    |
| `/user/:user` | `/user/you`         | ✓ match, `user = "you"`    |
| `/user/:user` | `/user/eko/profile` | ✗ no match (dua segmen)    |
| `/user/:user` | `/user/`            | ✗ no match (segmen kosong) |

```go
func TestHttprouterPatternNameParams(t *testing.T) {
	router := httprouter.New()

	// Route dengan dua named parameter: :id dan :itemsId.
	// Setiap `:nama` hanya menangkap satu segmen URL — tidak bisa lebih.
	// Pola ini cocok untuk resource bersarang seperti /products/1/items/5
	router.GET("/products/:id/items/:itemsId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")           // mengambil nilai segmen :id
		itemsId := p.ByName("itemsId") // mengambil nilai segmen :itemsId
		text := "Products: " + id + ", Item id: " + itemsId

		fmt.Fprint(w, text)
	})

	// URL /products/1/items/1 → id="1", itemsId="1"
	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Products: 1, Item id: 1", string(body))
}
```

## Catch-All Parameter

Catch-all parameter menangkap **semua sisa segmen** setelah posisinya, termasuk karakter `/`. Ditulis dengan awalan `*` dan harus berada di **akhir** URL pattern. Nilai yang ditangkap selalu diawali dengan `/`.

| Pattern          | URL                    | Hasil                                    |
| ---------------- | ---------------------- | ---------------------------------------- |
| `/src/*filepath` | `/src/`                | ✗ no match                               |
| `/src/*filepath` | `/src/somefile`        | ✓ match, `filepath = "/somefile"`        |
| `/src/*filepath` | `/src/subdir/somefile` | ✓ match, `filepath = "/subdir/somefile"` |

```go
func TestHttprouterPatternCatchAllParams(t *testing.T) {
	router := httprouter.New()

	// *image menangkap semua sisa path setelah /images.
	// Berbeda dengan named param, ini bisa menangkap beberapa segmen sekaligus.
	// Cocok digunakan untuk file serving atau nested path yang panjangnya tidak pasti.
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Nilai catch-all selalu diawali "/" — misalnya "/small/profile.png"
		image := p.ByName("image")
		text := "Image: " + image

		fmt.Fprint(w, text)
	})

	// /images/small/profile.png → image = "/small/profile.png"
	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Perhatikan nilai dimulai dengan "/" karena catch-all menyertakan slash awal
	assert.Equal(t, "Image: /small/profile.png", string(body))
}
```

## Cara Kerja

```
Named Parameter — menangkap satu segmen
──────────────────────────────────────────────────────
Route:   /products/:id/items/:itemsId
Request: /products/1/items/5
                   │         │
                   ▼         ▼
            id = "1"   itemsId = "5"

Setiap `:param` hanya cocok dengan satu segmen (tanpa slash).


Catch-All Parameter — menangkap sisa seluruh path
──────────────────────────────────────────────────────
Route:   /images/*image
Request: /images/small/profile.png
                 │
                 ▼
          image = "/small/profile.png"
          (semua sisa path ditangkap, termasuk slash)
```

## Perbandingan Named vs Catch-All

| Aspek                 | Named (`:id`)     | Catch-All (`*filepath`)   |
| --------------------- | ----------------- | ------------------------- |
| Simbol                | `:`               | `*`                       |
| Menangkap             | Satu segmen saja  | Semua sisa segmen         |
| Boleh di tengah route | ✓ Ya              | ✗ Harus di akhir          |
| Nilai mengandung `/`  | ✗ Tidak           | ✓ Ya (selalu diawali `/`) |
| Contoh use case       | ID resource, slug | File path, nested URL     |

## Alur Penggunaan

1. Tentukan apakah URL bersifat satu level (`/user/:id`) atau multi-level (`/files/*path`)
2. Gunakan `:nama` untuk menangkap satu segmen yang posisinya diketahui
3. Gunakan `*nama` di akhir route untuk menangkap sisa path yang tidak pasti panjangnya
4. Akses nilai keduanya dengan `p.ByName("nama")` di dalam handler
5. Ingat bahwa nilai catch-all selalu diawali `/` — strip jika tidak diinginkan

## Catatan Penting

| Catatan                             | Penjelasan                                                     |
| ----------------------------------- | -------------------------------------------------------------- |
| Catch-all harus di posisi terakhir  | Tidak bisa ada segmen lain setelah `*nama`                     |
| Nilai catch-all diawali `/`         | `/images/*img` dengan URL `/images/a/b` → `img = "/a/b"`       |
| Named param tidak mencocokkan slash | `:id` tidak akan menangkap `a/b` — gunakan catch-all untuk itu |
| Nama param harus unik per route     | Dua `:id` dalam satu route akan menyebabkan error              |

Next: [Serve File](./05-serve-file.md)
