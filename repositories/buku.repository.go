package repositories

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"perpustakaan/models"
)

const (
	CreateNewBook = `insert into bukus(id,kode_buku, penerbit, pengarang, jumlah_halaman, tahun_terbit, kategori_id, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	DeleteBookById = `delete from bukus where id=$1`
	GetAllWithPaginate = `select * from bukus offset $1 limit $2`
	UpdateStock = `update bukus set stock=$1 where id=$2`
	FindBookById = `select * from bukus where id = $1`
)

type IBukuRepository interface {
	CreateBuku(newBook *models.Buku) (*models.Buku, error)
	FindBookById(id uuid.UUID) (*models.Buku, error)
	GetAllBookByPaginate(page, size string) ([]*models.Buku, error)
	UpdateStock(book *models.Buku) (*models.Buku, error)
}

type bukuRepository struct {
	db *sql.DB
}

func (b *bukuRepository) FindBookById(id uuid.UUID) (*models.Buku, error) {
	book, err := b.db.Prepare(FindBookById)
	if err != nil {
		return nil, err
	}

	rows, errs := book.Query(id)
	if errs != nil {
		return nil, err
	}
	var buku models.Buku
	for rows.Next() {
		err = rows.Scan(&buku.ID, &buku.KodeBuku, &buku.Penerbit, &buku.Pengarang, &buku.JumlahHalaman, &buku.TahunTerbit, &buku.KategoriID, &buku.CreatedAt,&buku.UpdatedAt,&buku.DeletedAt,&buku.Stock)
		if err != nil {
			return nil, err
		}
	}
	return &buku, nil
}

func (b *bukuRepository) GetAllBookByPaginate(page, size string) ([]*models.Buku, error) {
	panic("implement me")
}

func (b *bukuRepository) UpdateStock(book *models.Buku) (*models.Buku, error) {
	_,err := b.db.Exec(UpdateStock, book.Stock,book.ID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *bukuRepository) CreateBuku(newBook *models.Buku) (*models.Buku, error) {

	_, err := b.db.Exec(CreateNewBook,newBook.ID, newBook.KodeBuku, newBook.Penerbit, newBook.Pengarang,newBook.JumlahHalaman, newBook.TahunTerbit, newBook.KategoriID,newBook.CreatedAt)
	if err != nil {
		return nil, err
	}

	return newBook, nil
}

func NewBukuRepository(db *sql.DB) IBukuRepository  {
	return &bukuRepository{db}
}
