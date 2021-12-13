package usecase

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)


type IPengembalianUsecase interface {
	CreatePengembalian(pengembalian *models.Pengembalian) (*models.Pengembalian, error)
}

type pengembalianUsecase struct {
	repo  repositories.IPengembalianRepo
}

func (p *pengembalianUsecase) CreatePengembalian(pengembalian *models.Pengembalian) (*models.Pengembalian, error) {
	return p.repo.CreatePengembalian(pengembalian)
}

func NewPengembalianUsecase(repo repositories.IPengembalianRepo) IPengembalianUsecase  {
	return &pengembalianUsecase{repo}
}
