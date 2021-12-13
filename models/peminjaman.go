package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type statusPinjam string

const (
	DIPINJAM statusPinjam = "Dipinjam"
	DIKEMBALIKAN statusPinjam = "Sudah Dikembalikan"
)
type Peminjaman struct {
	ID 				uuid.UUID		`json:"id"`
	TanggalPinjam 	string 			`json:"tanggal_pinjam"`
	LamaPinjam 		int			`json:"lama_pinjam"`
	Keterangan 		string			`json:"keterangan"`
	Status			statusPinjam	`json:"status"`
	MemberID 		uuid.UUID 		`json:"member_id" `
	CreatedAt   	time.Time  		`json:"created_at"`
	UpdatedAt   	*time.Time  		`json:"updated_at"`
	DeletedAt   	*time.Time 		`json:"deleted_at" sql:"index"`
}

func (p *Peminjaman) Prepare() error {
	p.ID = uuid.NewV4()
	p.CreatedAt = time.Now()
	return nil
}
