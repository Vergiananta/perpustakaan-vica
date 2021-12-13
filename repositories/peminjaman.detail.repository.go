package repositories

import (
	"database/sql"
	"perpustakaan/models"
)

const (
	INSERT = `insert into detail_pinjam (id, peminjaman_id, buku_id, quantity, created_at) values ($1,$2,$3,$4,$5)`
)

type IDetailPinjamRepo interface {
	CreateDetail(details *models.PeminjamanDetail) (*models.PeminjamanDetail, error)
}

type detailPinjamRepo struct {
	db *sql.DB
}

func (d *detailPinjamRepo) CreateDetail(details *models.PeminjamanDetail) (*models.PeminjamanDetail, error) {
	_,err := d.db.Exec(INSERT, details.ID,details.PeminjamanID, details.BukuID, details.Quantity, details.CreatedAt)
	if err != nil {
		return nil, err
	}
	return details, err
}

func NewDetailPinjamRepo(db *sql.DB) IDetailPinjamRepo {
	return &detailPinjamRepo{db}
}