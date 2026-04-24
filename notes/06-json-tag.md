# JSON Tag

## Gambaran Umum

| Konsep        | Detail                                                         |
| ------------- | -------------------------------------------------------------- |
| Masalah       | Nama field di struct (PascalCase) ≠ key di JSON (snake_case)   |
| Solusi        | JSON Tag pada struct menggunakan `json:"nama_key"`             |
| Mekanisme     | Tag Reflection — dibaca oleh package `json` saat encode/decode |
| Berlaku untuk | Encode (`Marshal`) dan Decode (`Unmarshal`)                    |

Secara default, `json.Marshal` dan `json.Unmarshal` mencocokkan field struct dengan key JSON berdasarkan **nama yang sama persis** (case-sensitive). Masalah muncul ketika konvensi penamaan berbeda — struct Go umumnya menggunakan **PascalCase**, sedangkan JSON API sering menggunakan **snake_case**.

Solusinya adalah **JSON Tag**: anotasi yang ditulis langsung di struct untuk menentukan nama key JSON yang digunakan saat konversi.

## Definisi Struct dengan JSON Tag

```go
type Product struct {
    ID       int     `json:"id"`        // field ID → key "id" di JSON
    Name     string  `json:"name"`      // field Name → key "name" di JSON
    Price    int     `json:"price"`     // field Price → key "price" di JSON  (case-insensitive saat decode)
    ImageURL string  `json:"image_url"` // field ImageURL → key "image_url" di JSON
}
```

Tanpa tag, `json.Marshal` akan menghasilkan key `"ID"`, `"Name"`, `"Price"`, `"ImageURL"`.
Dengan tag, key yang dihasilkan sesuai nilai tag: `"id"`, `"name"`, `"price"`, `"image_url"`.

## Contoh Kode — Decode dengan JSON Tag

```go
func TestJSONTagDecode(t *testing.T) {

    // JSON dari luar (API/server) menggunakan snake_case
    // "PRICE" tetap terbaca karena decode bersifat case-insensitive untuk tag
    jsonString := `{"id":1,"name":"Rinso","PRICE":1000,"image_url":""}`
    jsonBytes := []byte(jsonString)

    product := Product{}
    err := json.Unmarshal(jsonBytes, &product)

    if err != nil {
        panic(err)
    }

    // Field struct terisi sesuai pemetaan tag
    fmt.Println(product) // {1 Rinso 1000 }
}
```

## Cara Kerja

```
JSON masuk: {"id":1,"name":"Rinso","PRICE":1000,"image_url":""}
                │         │           │               │
                ▼         ▼           ▼               ▼
           tag:"id"  tag:"name"  tag:"price"   tag:"image_url"
                │         │           │               │
                ▼         ▼           ▼               ▼
Struct:      ID=1     Name="Rinso"  Price=1000   ImageURL=""
```

Saat **encode** (`Marshal`), proses ini berjalan terbalik — nama field struct dipetakan ke key JSON sesuai tag.

## Aturan Pencocokan Key saat Decode

| Kondisi                                        | Hasil                                   |
| ---------------------------------------------- | --------------------------------------- |
| Key JSON cocok dengan nilai tag (exact)        | ✅ Terbaca                              |
| Key JSON berbeda huruf kapital dari tag        | ✅ Tetap terbaca (case-insensitive)     |
| Key JSON tidak ada di tag manapun              | ⚠️ Diabaikan, field bernilai zero value |
| Tag ada tapi key sama sekali tidak ada di JSON | ⚠️ Field bernilai zero value            |
| Key mengandung `_` tapi tag tidak              | ❌ Tidak terbaca                        |

## Alur Penggunaan

1. **Definisikan struct** dengan tag `json:"nama_key"` pada setiap field
2. **Siapkan data JSON** — bisa dari string, file, atau response API
3. **Konversi ke `[]byte`** menggunakan `[]byte(jsonString)`
4. **Panggil `json.Unmarshal()`** dengan pointer ke struct
5. **Tangani error** — pastikan JSON valid dan tipe data kompatibel
6. **Gunakan struct** — field kini terisi sesuai pemetaan tag

## Catatan Penting

| Topik                            | Penjelasan                                                         |
| -------------------------------- | ------------------------------------------------------------------ |
| Tag wajib diapit backtick        | Penulisan: `` `json:"nama_key"` `` — bukan tanda kutip biasa       |
| `json:"-"`                       | Field dengan tag ini akan **dilewati** saat encode maupun decode   |
| `json:",omitempty"`              | Field dengan zero value tidak akan disertakan saat encode          |
| `json:"nama,omitempty"`          | Gabungan nama custom dan omitempty                                 |
| Decode bersifat case-insensitive | `"PRICE"` dapat cocok dengan tag `"price"` saat Unmarshal          |
| Encode bersifat exact            | Saat Marshal, key output **selalu** sesuai nilai tag, tidak diubah |

Next: [Map](./07-map.md)
