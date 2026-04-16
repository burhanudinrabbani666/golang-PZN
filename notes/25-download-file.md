# Download file

## Download File

- Selain upload file, kadang kita ingin membuat halaman website yang digunakan untuk download sesuatu
- Sebenarnya di Go-Lang sudah disediakan menggunakan FileServer dan ServeFile
- Dan jika kita ingin memaksa file di download (tanpa di render oleh browser, kita bisa menggunakan header Content-Disposition)
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition

```go
func DownloadFile(writer http.ResponseWriter, request *http.Request) {

	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, http.StatusBadRequest)

		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    localhost,
		Handler: http.HandlerFunc(DownloadFile),
	}

	errorDownload := server.ListenAndServe()

	if errorDownload != nil {
		panic(errorDownload)
	}

}
```

Next: []()
