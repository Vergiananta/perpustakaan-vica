package repositories

import (
	"database/sql"
	"perpustakaan/models"
)

const (
	INSERTPINJAM = `insert into pinjam (id, tanggal_pinjam, lama_pinjam, keterangan, status, member_id, created_at) values ($1,$2,$3,$4,$5,$6,$7)`
)

type IPeminjamanRepo interface {
	CreatePinjam(pinjam *models.Peminjaman) (*models.Peminjaman, error)
}

type peminjamanRepo struct {
	db *sql.DB
}

func (p *peminjamanRepo) CreatePinjam(pinjam *models.Peminjaman) (*models.Peminjaman, error) {
	_,err:= p.db.Exec(INSERTPINJAM, pinjam.ID,pinjam.TanggalPinjam,pinjam.LamaPinjam, pinjam.Keterangan,pinjam.Status, pinjam.MemberID,pinjam.CreatedAt)
	if err != nil {
		return nil, err
	}
	return pinjam, nil
}

func NewPeminjamanRepo(db *sql.DB) IPeminjamanRepo {
	return &peminjamanRepo{db}
}