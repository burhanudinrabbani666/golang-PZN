package golangpzn

import (
	"fmt"
	"net/http"
	"testing"
)

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
