package main

import (
	"github.com/aadi-1024/gotasks-web/pkg/db"
	"github.com/aadi-1024/gotasks-web/pkg/handlers"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

func main() {
	dbRepo, err := db.NewDbRepo("host=localhost port=5432 user=postgres password=password database=tasks")
	handlerRepo := handlers.NewRepository(dbRepo)

	if err != nil {
		log.Fatalln(err)
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: NewRouter(handlerRepo),
	}
	srv.ListenAndServe()
}
