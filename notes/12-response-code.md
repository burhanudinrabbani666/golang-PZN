# Response code

- Dalam HTTP, terdapat yang namanya response code
- Response code merupakan representasi kode response
- Dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses oleh server atau gagal
- Ada banyak sekali response code yang bisa kita gunakan saat membuat web
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

## Mengubah Response Code

- Secara default, jika kita tidak menyebutkan response code, maka response code nya adalah 200 OK
- Jika kita ingin mengubahnya, kita bisa menggunakan function ResponseWriter.WriteHeader(int)
- Semua data status code juga sudah disediakan di Go-Lang, jadi kita ingin, kita bisa gunakan variable yang sudah disediakan : https://github.com/golang/go/blob/master/src/net/http/status.go

```go
func ResponseCode(w http.ResponseWriter, r http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is Emplty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

// Not Valid
func TestResCode(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, *req)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func TestResCodeValid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/?name=Burhanudin", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, *req)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}
```

Next: [Cookie](./13-cookie.md)
