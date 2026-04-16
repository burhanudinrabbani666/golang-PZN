# Upload file

- Saat membuat web, selain menerima input data berupa form dan query param, kadang kita juga menerima input data berupa file dari client
- Go-Lang Web sudah memiliki fitur untuk management upload file
- Hal ini memudahkan kita ketika butuh membuat web yang menerima input file upload

## MultiPart

- Saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan Request.FormFile(name), di dalam nya secara otomatis melakukan parsing terlebih dahulu
- Hasilnya merupakan data-data yang terdapat pada package multipart, seperti multipart.File sebagai representasi file nya, dan multipart.FileHeader sebagai informasi file nya

```go
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	MyTemplate.ExecuteTemplate(writer, "upload_form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {

	file, fileHeader, errorForm := request.FormFile("file")
	if errorForm != nil {
		panic(errorForm)
	}

	fileDestinations, errorCreate := os.Create("./resources/" + fileHeader.Filename)
	if errorCreate != nil {
		panic(errorCreate)
	}

	_, errorCopy := io.Copy(fileDestinations, file)
	if errorCopy != nil {
		panic(errorCopy)
	}

	name := request.PostFormValue("name")
	MyTemplate.ExecuteTemplate(writer, "upload_success.gohtml", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    localhost,
		Handler: mux,
	}

	errorUpload := server.ListenAndServe()

	if errorUpload != nil {
		panic(errorUpload)
	}

}

//go:embed resources/wallpaper5.jpg
var UploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Burhanudin Rabbani")
	file, _ := writer.CreateFormFile("file", "Contoh Upload .png")
	file.Write(UploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, localhostFull+"/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))

}
```

Next: [Download file](./25-download-file.md)
