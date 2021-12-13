package usecase

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type IPeminjamanUsecase interface {
	CreatePinjam(pinjam *models.Peminjaman) (*models.Peminjaman, error)
}

type peminjamanUsecase struct {
	repo repositories.IPeminjamanRepo
}

func (p *peminjamanUsecase) CreatePinjam(pinjam *models.Peminjaman) (*models.Peminjaman, error) {
	return p.repo.CreatePinjam(pinjam)
}

func NewPeminjamanUsecase(repo repositories.IPeminjamanRepo) IPeminjamanUsecase {
	return &peminjamanUsecase{repo}
}