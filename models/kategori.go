package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Kategori struct {
	ID 			uuid.UUID	`json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt   time.Time  	`json:"created_at"`
	UpdatedAt   *time.Time  	`json:"updated_at"`
	DeletedAt   *time.Time 	`json:"deleted_at" sql:"index"`
}

func (k *Kategori) Prepare()  {
	k.ID = uuid.NewV4()
	k.CreatedAt = time.Now()
}