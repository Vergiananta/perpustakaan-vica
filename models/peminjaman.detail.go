package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type PeminjamanDetail struct {
	ID 				uuid.UUID	`json:"id"`
	PeminjamanID 	uuid.UUID	`json:"peminjaman_id"`
	BukuID 			uuid.UUID	`json:"buku_id"`
	Quantity		int32		`json:"quantity"`
	CreatedAt   	time.Time  	`json:"created_at"`
	UpdatedAt   	*time.Time  	`json:"updated_at"`
	DeletedAt   	*time.Time 	`json:"deleted_at" sql:"index"`
}

func (d *PeminjamanDetail) Prepare() error {
	d.ID = uuid.NewV4()
	d.CreatedAt = time.Now()
	return nil
}

