# Streaming Decoder

## Gambaran Umum

| Konsep                     | Detail                                                       |
| -------------------------- | ------------------------------------------------------------ |
| Fungsi utama               | `json.NewDecoder(reader)` + `.Decode(&target)`               |
| Sumber data                | `io.Reader` — file, network, HTTP request body, dsb          |
| Perbedaan dari `Unmarshal` | Tidak perlu membaca seluruh data ke memory terlebih dahulu   |
| Keuntungan                 | Lebih efisien untuk data besar atau data yang terus mengalir |

Sebelumnya kita menggunakan `json.Unmarshal()` yang mengharuskan data JSON sudah tersimpan dalam variabel `[]byte`. Pendekatan ini tidak masalah untuk data kecil, namun kurang efisien jika sumber datanya adalah **stream** seperti file besar, koneksi jaringan, atau HTTP request body.

`json.NewDecoder()` memungkinkan kita membaca dan mengonversi JSON **langsung dari stream** tanpa harus memuat seluruh isi ke memory terlebih dahulu.

## Contoh Kode

```go
func TestJsonStreamingDecoder(t *testing.T) {

    // Buka file JSON sebagai io.Reader
    // os.Open() mengembalikan *os.File yang mengimplementasikan io.Reader
    reader, err := os.Open("./data/customer.json")
    if err != nil {
        panic(err)
    }
    defer reader.Close() // pastikan file ditutup setelah selesai digunakan

    // Buat decoder yang membaca langsung dari reader (stream)
    // Tidak ada data yang dibaca ke memory di tahap ini
    decoder := json.NewDecoder(reader)

    // Siapkan slice sebagai wadah hasil decode
    // Sesuaikan tipe dengan struktur JSON yang dibaca
    customer := []Customer{}

    // Baca stream dan decode JSON langsung ke dalam customer
    err = decoder.Decode(&customer)
    if err != nil {
        panic(err)
    }

    fmt.Println(customer) // [{Budi D Santoso} {Siti A Rahma} ...]
}
```

## Isi File `customer.json`

```json
[
  { "FirstName": "Budi", "MiddleName": "D", "LastName": "Santoso" },
  { "FirstName": "Siti", "MiddleName": "A", "LastName": "Rahma" }
]
```

## Cara Kerja

```
  File / Network / Request Body
            │
            │  io.Reader (stream — data belum dibaca)
            ▼
   json.NewDecoder(reader)
            │
            │  Decoder siap, stream belum diproses
            ▼
    decoder.Decode(&customer)
            │
            ├─ Baca data dari stream secara bertahap
            ├─ Parse JSON setiap kali data tersedia
            └─ Isi field struct/slice sesuai key JSON
                        │
                        ▼
            customer = [{Budi D Santoso} {Siti A Rahma}]
```

## Perbandingan: `Unmarshal` vs `Decoder`

| Aspek                | `json.Unmarshal`           | `json.NewDecoder`              |
| -------------------- | -------------------------- | ------------------------------ |
| Sumber data          | `[]byte` (sudah di memory) | `io.Reader` (stream)           |
| Cocok untuk          | Data kecil, sudah tersedia | File besar, network, HTTP body |
| Penggunaan memory    | Seluruh data dimuat dulu   | Dibaca bertahap                |
| Kemudahan            | Lebih sederhana            | Sedikit lebih panjang          |
| Multiple JSON object | ❌ Tidak bisa langsung     | ✅ Bisa dengan loop `Decode()` |

## Alur Penggunaan

1. **Dapatkan `io.Reader`** — dari `os.Open()`, HTTP response body, atau sumber lain
2. **Buat decoder** — `json.NewDecoder(reader)`
3. **Siapkan target** — struct, slice, atau `map[string]any`
4. **Panggil `Decode()`** — masukkan pointer ke target
5. **Tangani error** — periksa nilai error yang dikembalikan
6. **Tutup reader** — gunakan `defer reader.Close()` untuk file

## Catatan Penting

| Topik                  | Penjelasan                                                                                 |
| ---------------------- | ------------------------------------------------------------------------------------------ |
| `defer reader.Close()` | Selalu tutup file atau koneksi setelah selesai untuk mencegah resource leak                |
| `io.Reader`            | Interface yang diimplementasikan oleh `*os.File`, `http.Request.Body`, `bytes.Reader`, dsb |
| Multiple decode        | Memanggil `Decode()` berkali-kali akan membaca JSON object berikutnya dari stream          |
| Error `io.EOF`         | Dikembalikan saat stream sudah habis dibaca — bukan error fatal                            |

Next: [Streaming Encoder](./09-streaming-encoder.md)
