package models

import (
	"fmt"
	"time"
)

type TaskDataModel struct {
	Id     int
	Title  string
	Expiry time.Time
}

func (t *TaskDataModel) String() string {
	return fmt.Sprintf("%v %v %v", t.Id, t.Title, t.Expiry)
}
