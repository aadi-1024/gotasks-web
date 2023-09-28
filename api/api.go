package api

import (
	"net/http"
	"os"
)

var htmlTemplates map[string][]byte

func loadTemplates() error {
	htmlTemplates = make(map[string][]byte)
	content, err := os.ReadFile("templates/index.html")
	htmlTemplates["index"] = content
	return err

}

func AddRoutes(handler *http.ServeMux) error {

	err := loadTemplates()
	if err != nil {
		return err
	}

	//index
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("content-type", "text/html")
		writer.Header().Add("connection", "close")
		writer.WriteHeader(200)
		_, _ = writer.Write(htmlTemplates["index"])
	})

	return nil
}
