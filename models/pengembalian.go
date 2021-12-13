package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Pengembalian struct {
	ID 				uuid.UUID	`json:"id" `
	PeminjamanID 	uuid.UUID	`json:"peminjaman_id"`
	TanggalKembali	string 		`json:"tanggal_kembali"`
	CreatedAt   	time.Time  	`json:"created_at"`
	UpdatedAt   	*time.Time  `json:"updated_at"`
	DeletedAt   	*time.Time 	`json:"deleted_at" sql:"index"`
}

func (p *Pengembalian) Prepare() error {
	p.ID = uuid.NewV4()
	p.CreatedAt = time.Now()
	return nil
}