package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aadi-1024/gotasks-web/pkg/models"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"time"
)

type Repository struct {
	conn *sql.DB
}

const (
	maxIdleConnections = 5
	maxWorkConnections = 10
	maxIdleTime        = time.Minute
	dbTimeout          = 3 * time.Second
)

func NewDbRepo(dsn string) (*Repository, error) {
	conn, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(maxIdleConnections)
	conn.SetConnMaxIdleTime(maxIdleTime)
	conn.SetMaxOpenConns(maxWorkConnections)

	err = conn.Ping()
	if err != nil {
		return &Repository{conn}, errors.New(fmt.Sprintf("couldn't ping database: %v", err))
	}
	return &Repository{conn}, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*models.TaskDataModel, error) {
	query := `select id, title, expiry from tasks`
	var model *models.TaskDataModel
	buffer := make([]*models.TaskDataModel, 0)

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while querying: %v", err))
	}
	defer rows.Close()

	for rows.Next() {
		model = &models.TaskDataModel{}
		err = rows.Scan(
			&model.Id,
			&model.Title,
			&model.Expiry,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		buffer = append(buffer, model)
	}

	if err = rows.Err(); err != nil {
		return buffer, err
	}
	return buffer, nil
}

func (r *Repository) GetById(ctx context.Context, id string) (*models.TaskDataModel, error) {
	data := &models.TaskDataModel{}
	query := `select id, title, expiry from tasks where id=$1`

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	row := r.conn.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&data.Id,
		&data.Title,
		&data.Expiry,
	)

	return data, err
}

func (r *Repository) DeleteId(ctx context.Context, id string) error {
	query := `delete from tasks where id=$1`

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	_, err := r.conn.QueryContext(ctx, query, id)

	if err != nil {
		log.Println("failed deleting from database: ", err)
	}
	return err
}

func (r *Repository) UpdateTask(ctx context.Context, data *models.TaskDataModel) error {
	query := `update tasks set title=$1, expiry=$2 where id=$3`

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	_, err := r.conn.QueryContext(ctx, query, data.Title, data.Expiry, data.Id)

	if err != nil {
		log.Println("failed updating in database: ", err)
	}
	return err
}
