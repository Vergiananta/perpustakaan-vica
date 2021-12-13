package repositories

import (
	"database/sql"
	"perpustakaan/models"
)

const (
	INSERTKATEGORI = `insert into kategori (id, name, created_at) values ($1,$2,$3)`
)

type IKategoriRepo interface {
	CreateCategory(kategori *models.Kategori) (*models.Kategori, error)
}

type kategoriRepo struct {
	db *sql.DB
}

func (k *kategoriRepo) CreateCategory(kategori *models.Kategori) (*models.Kategori, error) {
	_,err := k.db.Exec(INSERTKATEGORI, kategori.ID,kategori.Name,kategori.CreatedAt)
	if err != nil {
		return nil,err
	}
	return kategori, nil
}

func NewKategoriRepo(db *sql.DB) IKategoriRepo {
	return &kategoriRepo{db}
}