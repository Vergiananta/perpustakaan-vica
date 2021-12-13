package usecase

import (
	uuid "github.com/satori/go.uuid"
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type IBukuUsecase interface {
	CreateBook(newBuku *models.Buku) (*models.Buku, error)
	FindBookById(id uuid.UUID) (*models.Buku, error)
	UpdateStock(book *models.Buku) (*models.Buku, error)
}

type bukuUsecase struct {
	bukuRepo repositories.IBukuRepository
}

func (b *bukuUsecase) UpdateStock(book *models.Buku) (*models.Buku, error) {
	return b.bukuRepo.UpdateStock(book)
}

func (b *bukuUsecase) FindBookById(id uuid.UUID) (*models.Buku, error) {
	return b.bukuRepo.FindBookById(id)
}

func (b *bukuUsecase) CreateBook(newBuku *models.Buku) (*models.Buku, error) {
	newBuku.Prepare()
	return b.bukuRepo.CreateBuku(newBuku)
}

func NewBukuUsecase(bukuRepo repositories.IBukuRepository) IBukuUsecase  {
	return &bukuUsecase{bukuRepo}
}


