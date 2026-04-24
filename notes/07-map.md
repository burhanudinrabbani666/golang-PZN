# JSON dengan Map

## Gambaran Umum

| Konsep         | Detail                                                    |
| -------------- | --------------------------------------------------------- |
| Tipe data      | `map[string]any` (alias dari `map[string]interface{}`)    |
| Digunakan saat | Struktur JSON tidak diketahui atau berubah-ubah (dynamic) |
| Key map        | Berasal dari atribut/key JSON                             |
| Value map      | Bertipe `any` — perlu type assertion untuk digunakan      |
| JSON Tag       | ❌ Tidak didukung pada tipe Map                           |

Saat struktur JSON sudah diketahui dan tetap, kita gunakan **struct**. Namun ketika atribut JSON bisa bertambah, berkurang, atau tidak menentu, menggunakan struct menjadi tidak praktis karena semua field harus dideklarasikan terlebih dahulu.

Solusinya adalah `map[string]any` — Go akan secara otomatis memetakan setiap key JSON menjadi key map, dan setiap value JSON menjadi value map bertipe `any`.

## Contoh Kode — Decode JSON ke Map

```go
// Nama fungsi asli di kode sumber: TestJSONMapEncode (keliru)
// Operasi yang dilakukan adalah Unmarshal = Decode
func TestJSONMapDecode(t *testing.T) {

    // JSON dengan struktur yang bisa saja berubah-ubah
    jsonString := `{"id":"AP-001","name":"Apple Macbook Pro","price":1234,"image_url":"https://example.com"}`
    jsonByte := []byte(jsonString)

    // map[string]any: key bertipe string, value bertipe any (interface{})
    var result map[string]any
    json.Unmarshal(jsonByte, &result)

    // Mengakses value seperti mengakses map biasa
    fmt.Println(result)           // map[id:AP-001 image_url:https://example.com name:Apple Macbook Pro price:1234]
    fmt.Println(result["id"])     // AP-001
    fmt.Println(result["name"])   // Apple Macbook Pro
}
```

## Contoh Kode — Encode Map ke JSON

```go
// Nama fungsi asli di kode sumber: TestJSONMapDecode (keliru)
// Operasi yang dilakukan adalah Marshal = Encode
func TestJSONMapEncode(t *testing.T) {

    // Buat map dengan key-value bebas tanpa perlu mendefinisikan struct
    product := map[string]any{
        "id":    "P-001",
        "name":  "Apple Mac Book",
        "price": 1234,
    }

    // Marshal mengonversi map menjadi JSON bytes
    bytes, _ := json.Marshal(product)

    fmt.Println(string(bytes)) // {"id":"P-001","name":"Apple Mac Book","price":1234}
}
```

## Cara Kerja

```
         DECODE (Unmarshal)
─────────────────────────────────────────
JSON:  {"id":"AP-001","name":"Rinso","price":1000}
                │
                ▼
       map[string]any{
           "id"    → "AP-001"   (string)
           "name"  → "Rinso"    (string)
           "price" → 1000       (float64)  ← angka JSON selalu jadi float64!
       }

         ENCODE (Marshal)
─────────────────────────────────────────
       map[string]any{
           "id"    → "P-001"
           "name"  → "Apple Mac Book"
           "price" → 1234
       }
                │
                ▼
JSON:  {"id":"P-001","name":"Apple Mac Book","price":1234}
```

## Type Assertion saat Mengakses Value

Karena value bertipe `any`, kita perlu **type assertion** untuk mendapatkan tipe aslinya:

```go
// Ambil value dan lakukan type assertion
id := result["id"].(string)       // assert ke string
price := result["price"].(float64) // angka JSON selalu float64, bukan int!

fmt.Println(id)    // AP-001
fmt.Println(price) // 1234
```

## Perbandingan: Struct vs Map

| Aspek           | Struct                  | `map[string]any`        |
| --------------- | ----------------------- | ----------------------- |
| Struktur JSON   | Tetap / diketahui       | Dinamis / tidak menentu |
| JSON Tag        | ✅ Didukung             | ❌ Tidak didukung       |
| Type safety     | ✅ Langsung sesuai tipe | ⚠️ Perlu type assertion |
| Akses field     | `product.Name`          | `result["name"]`        |
| Tipe angka JSON | Sesuai tipe field       | Selalu `float64`        |

## Catatan Penting

| Topik                     | Penjelasan                                                                        |
| ------------------------- | --------------------------------------------------------------------------------- |
| Angka JSON → `float64`    | Semua angka dalam JSON otomatis menjadi `float64` saat decode ke `map[string]any` |
| Type assertion bisa panic | Gunakan bentuk `v, ok := result["key"].(string)` untuk menghindari panic          |
| Urutan key saat encode    | Key map tidak dijamin urutannya saat di-encode ke JSON                            |
| JSON Tag tidak berlaku    | Tag `json:"..."` diabaikan pada tipe map                                          |

Next: [Streaming Decoder](./08-streaming-decoder.md)
