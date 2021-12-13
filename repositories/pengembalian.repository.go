package repositories

import (
	"database/sql"
	"perpustakaan/models"
)

const (
	INSERTPENGEMBALIAN = `insert into return_book (id, peminjaman_id, tanggal_kembali, created_at) values ($1,$2,$3,$4)`
)

type IPengembalianRepo interface {
	CreatePengembalian(pengembalian *models.Pengembalian) (*models.Pengembalian, error)
}

type pengembalianRepo struct {
	db *sql.DB
}

func (p *pengembalianRepo) CreatePengembalian(pengembalian *models.Pengembalian) (*models.Pengembalian, error) {
	_,err := p.db.Exec(INSERTPENGEMBALIAN, pengembalian.ID,pengembalian.PeminjamanID,pengembalian.TanggalKembali,pengembalian.CreatedAt)
	if err != nil {
		return nil, err
	}
	return pengembalian, err
}

func NewPengembalianRepo(db *sql.DB) IPengembalianRepo {
	return &pengembalianRepo{db}
}
