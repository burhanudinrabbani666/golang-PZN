# JSON Array

## Gambaran Umum

| Konsep                        | Detail                                               |
| ----------------------------- | ---------------------------------------------------- |
| Representasi JSON Array di Go | Tipe data `slice` (`[]Type`)                         |
| Encode slice → JSON Array     | `json.Marshal(slice)`                                |
| Decode JSON Array → slice     | `json.Unmarshal(bytes, &slice)`                      |
| Isi JSON Array                | Bisa primitif (`string`, `int`) atau object kompleks |

JSON Array di Go direpresentasikan sebagai **slice**. Package `json` secara otomatis menangani konversi antara JSON Array dan slice Go — baik encode maupun decode.

## 1. Encode: Struct dengan Field Slice

Field `Hobbies` bertipe `[]string` akan otomatis dikonversi menjadi JSON Array.

```go
func TestJSONArray(t *testing.T) {

    customer := Customer{
        FirstName:  "Burhanudin",
        MiddleName: "D",
        LastName:   "Rabbani",
        Hobbies:    []string{"Gaming", "Coding", "Reading"},
    }

    bytes, _ := json.Marshal(customer)
    fmt.Println(string(bytes))

    // Output:
    // {"FirstName":"Burhanudin","MiddleName":"D","LastName":"Rabbani",
    //  "Hobbies":["Gaming","Coding","Reading"]}
    //
    // ✅ Slice []string → JSON Array ["Gaming","Coding","Reading"]
}
```

## 2. Decode: JSON Array ke Field Slice

JSON Array dalam key `"Hobbies"` otomatis diisi ke field `Hobbies []string` di struct.

```go
func TestJSONArrayDecode(t *testing.T) {

    jsonString := `{"FirstName":"Burhanudin","MiddleName":"D","LastName":"Rabbani","Hobbies":["Gaming","Coding","Reading"]}`
    jsonBytes := []byte(jsonString)

    customer := &Customer{}

    err := json.Unmarshal(jsonBytes, customer)
    if err != nil {
        panic(err)
    }

    fmt.Println(customer.Hobbies) // [Gaming Coding Reading]
    // ✅ JSON Array → kembali menjadi slice []string
}
```

## 3. Encode: Object Kompleks dengan Nested Array

Struct bisa memiliki field slice of struct (`[]Address`), yang akan dikonversi menjadi JSON Array of Object.

```go
func TestJSONArrayComplexEncode(t *testing.T) {

    customer := Customer{
        FirstName: "Burhanudin",
        Age:       23,
        IsMarried: false,
        Hobbies:   []string{"Gaming", "Coding", "Reading"},
        Addresses: []Address{    // []Address → JSON Array of Object
            {Street: "Jalan Belum Ada",     Country: "Indonesia", Postal: "12344"},
            {Street: "Jalan lagi dibangun", Country: "Malaysia",  Postal: "23456"},
        },
    }

    bytes, err := json.Marshal(customer)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println(string(bytes))
    // "Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","Postal":"12344"},
    //              {"Street":"Jalan lagi dibangun","Country":"Malaysia","Postal":"23456"}]
}
```

## 4. Decode: Root JSON Array (Tanpa Object Wrapper)

JSON tidak selalu dimulai dari `{}`. Jika root-nya langsung `[...]`, gunakan slice sebagai target decode.

```go
func TestOnlyJSONArray(t *testing.T) {

    // JSON yang dimulai langsung dari array — tanpa wrapper object
    jsonString := `[
        {"Street":"Jalan Belum Ada",     "Country":"Indonesia", "Postal":"12344"},
        {"Street":"Jalan lagi dibangun", "Country":"Malaysia",  "Postal":"23456"}
    ]`

    jsonBytes := []byte(jsonString)
    addresses := []Address{} // Target adalah slice, bukan struct tunggal

    err := json.Unmarshal(jsonBytes, &addresses)
    if err != nil {
        panic(err)
    }

    for index, address := range addresses {
        fmt.Printf("Index ke : %d\n", index)
        fmt.Printf("Street   : %s\n", address.Street)
        fmt.Printf("Country  : %s\n", address.Country)
        fmt.Printf("Postal   : %s\n", address.Postal)
    }
}
```

## Cara Kerja

```
ENCODE (Go → JSON)

  []string{"Gaming","Coding","Reading"}  →  ["Gaming","Coding","Reading"]
  []Address{{...},{...}}                 →  [{"Street":"..."},{"Street":"..."}]


DECODE (JSON → Go)

  JSON Array of string  ["Gaming","Coding","Reading"]  →  []string
  JSON Array of object  [{"Street":"..."},...]         →  []Address
  Root JSON Array       [...]                          →  []Address (bukan &Customer{})
```

## Alur Penggunaan

1. **Tentukan tipe array** — apakah isi array primitif (`[]string`) atau object (`[]Address`)?
2. **Definisikan field struct** sebagai slice dari tipe yang sesuai
3. **Encode** — `json.Marshal()` mengkonversi slice menjadi JSON Array secara otomatis
4. **Decode** — `json.Unmarshal()` mengisi slice dari JSON Array secara otomatis
5. **Root array** — jika JSON dimulai dari `[...]`, gunakan slice sebagai target, bukan struct

## Catatan Penting

| Topik               | Penjelasan                                                                    |
| ------------------- | ----------------------------------------------------------------------------- |
| Slice kosong vs nil | `[]Address{}` menghasilkan `[]` di JSON; slice `nil` menghasilkan `null`      |
| Root array          | Jika JSON root adalah `[...]`, target decode harus slice, bukan struct        |
| Nested array        | `[]Address` di dalam struct akan menjadi JSON Array of Object secara otomatis |
| Urutan elemen       | Urutan elemen dalam slice dipertahankan setelah encode maupun decode          |

Next: [JSON Tag](./06-json-tag.md)
