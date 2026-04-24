# Decode JSON

## Gambaran Umum

| Konsep            | Detail                                |
| ----------------- | ------------------------------------- |
| Fungsi utama      | `json.Unmarshal([]byte, interface{})` |
| Arah konversi     | JSON (`[]byte`) → Struct Go           |
| Parameter pertama | Data JSON dalam bentuk `[]byte`       |
| Parameter kedua   | Pointer ke struct tujuan (`&struct`)  |
| Return value      | `error` (nil jika berhasil)           |

Sebelumnya kita telah belajar **encode** — mengubah struct Go menjadi JSON. Sekarang kita belajar kebalikannya: **decode**, yaitu mengubah data JSON kembali menjadi struct Go menggunakan `json.Unmarshal()`.

## Contoh Kode

```go
func TestDecodeJSON(t *testing.T) {

    // JSON string mentah yang akan kita decode
    // Key harus sesuai dengan nama field di struct Customer
    jsonString := `{
        "FirstName":"Burhanudin",
        "MiddleName":"D",
        "LastName":"Rabbani"
    }`

    // Konversi string ke []byte karena json.Unmarshal menerima []byte
    // Logika: json.Marshal() → menghasilkan []byte → dikonversi ke string
    //         json.Unmarshal() ← menerima []byte ← konversi balik dari string
    jsonBytes := []byte(jsonString)

    // Siapkan struct kosong sebagai wadah hasil decode
    customer := Customer{}

    // Unmarshal mengisi field-field customer sesuai key yang ada di JSON
    // Gunakan pointer (&customer) agar perubahan berlaku pada variabel asli
    err := json.Unmarshal(jsonBytes, &customer)

    if err != nil {
        panic(err)
    }

    // Struct customer kini telah terisi dari data JSON
    fmt.Println(customer)            // {Burhanudin D Rabbani}
    fmt.Println(customer.FirstName)  // Burhanudin
    fmt.Println(customer.MiddleName) // D
    fmt.Println(customer.LastName)   // Rabbani
}
```

## Cara Kerja

```
JSON String (teks biasa)
        │
        ▼
  []byte(jsonString)        ← Konversi manual ke bytes
        │
        ▼
 json.Unmarshal(bytes, &customer)
        │
        ├─ Baca setiap key di JSON  →  Cari field yang cocok di struct
        │
        └─ Isi nilai field struct sesuai value di JSON
                │
                ▼
         customer.FirstName  = "Burhanudin"
         customer.MiddleName = "D"
         customer.LastName   = "Rabbani"
```

## Alur Penggunaan

1. **Siapkan JSON** — berupa string atau data yang diterima dari API/file
2. **Konversi ke `[]byte`** — gunakan `[]byte(jsonString)`
3. **Siapkan struct kosong** — sebagai wadah hasil decode
4. **Panggil `json.Unmarshal()`** — masukkan bytes dan pointer ke struct
5. **Tangani error** — selalu periksa nilai error yang dikembalikan
6. **Gunakan struct** — field-field kini sudah terisi dari data JSON

## Catatan Penting

| Topik                   | Penjelasan                                                                            |
| ----------------------- | ------------------------------------------------------------------------------------- |
| Kenapa pointer?         | `json.Unmarshal` perlu memodifikasi struct asli, bukan salinannya                     |
| Pencocokan key          | Key JSON harus **sama persis** (case-sensitive) dengan nama field struct              |
| Field tidak cocok       | Jika key di JSON tidak ada di struct, field tersebut diabaikan                        |
| Field tidak ada di JSON | Field struct yang tidak ada di JSON akan tetap bernilai zero value                    |
| Tipe data               | Nilai dalam JSON harus kompatibel dengan tipe field di struct                         |
| String vs non-string    | Nilai JSON yang diapit `""` adalah string; tanpa `""` bisa berupa number atau boolean |

Next: [JSON Array](./05-json-array.md)
