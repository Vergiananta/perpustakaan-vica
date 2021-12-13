package usecase

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type IKategoriUsecase interface {
	CreateKategori(category *models.Kategori) (*models.Kategori, error)
}

type kategoriUsecase struct {
	repo repositories.IKategoriRepo
}

func (k *kategoriUsecase) CreateKategori(category *models.Kategori) (*models.Kategori, error) {
	category.Prepare()
	return k.repo.CreateCategory(category)
}

func NewKategoriUsecase(repo repositories.IKategoriRepo) IKategoriUsecase {
	return &kategoriUsecase{repo}
}
