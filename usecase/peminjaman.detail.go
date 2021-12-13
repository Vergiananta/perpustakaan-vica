package usecase

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type IDetailPinjam interface {
	CreateDetailPinjam(details *models.PeminjamanDetail) (*models.PeminjamanDetail, error)
}

type detailPinjam struct {
	repo 	repositories.IDetailPinjamRepo
}

func (d *detailPinjam) CreateDetailPinjam(details *models.PeminjamanDetail) (*models.PeminjamanDetail, error) {
	return d.repo.CreateDetail(details)
}

func NewDetailPinjam(repo repositories.IDetailPinjamRepo) IDetailPinjam {
	return &detailPinjam{repo}
}
