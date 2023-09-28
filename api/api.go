package api

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Entry struct {
	Id   int
	Name string
	Time string
}

type PageData struct {
	PageTitle string
	Entries   []Entry
}

var tmpl *template.Template

func loadTemplates() error {

	tmpl = template.New("index")

	content, err := os.ReadFile("templates/index.html")

	tmpl = template.Must(tmpl.Parse(string(content)))
	//test data

	return err

}

func AddRoutes(handler *http.ServeMux) error {

	err := loadTemplates()
	if err != nil {
		return err
	}

	pageData := PageData{
		PageTitle: "Ayo this works",
		Entries: []Entry{
			Entry{
				Id:   0,
				Name: "Do this",
				Time: "12:00",
			},
			Entry{
				Id:   1,
				Name: "Do that",
				Time: "1:00",
			},
		},
	}

	//index
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("content-type", "text/html")
		writer.Header().Add("connection", "close")
		writer.WriteHeader(200)
	})

	handler.HandleFunc("/task", func(writer http.ResponseWriter, request *http.Request) {

		switch request.Method {
		case "GET":
			//writer.Header().Add("content-type", "text/html")
			//writer.Header().Add("connection", "close")
			//writer.WriteHeader(200)
			//_, _ = writer.Write([]byte(fmt.Sprintf("<h1>%v</h1>", request.URL.Query()["id"])))
			err = tmpl.Execute(writer, pageData)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	return nil
}
