package handlers

import (
	"bytes"
	"context"
	"github.com/aadi-1024/gotasks-web/pkg/db"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Repository struct {
	DbConn *db.Repository
}

func NewRepository(Db *db.Repository) *Repository {
	return &Repository{
		Db,
	}
}

func (m *Repository) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	data, err := m.DbConn.GetAll(context.Background())
	if err != nil {
		log.Println("error fetching data: ", err)
		http.Error(w, err.Error(), 400)
	}

	buf := new(bytes.Buffer)

	for _, val := range data {
		_, err = buf.Write([]byte(val.String()))
		if err != nil {
			log.Println(err)
		}
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := m.DbConn.GetById(context.Background(), id)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write([]byte(data.String()))
}
