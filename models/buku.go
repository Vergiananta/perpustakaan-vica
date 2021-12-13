package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Buku struct {
	ID 				uuid.UUID	`json:"id"`
	KodeBuku 		string		`json:"kode_buku"`
	Penerbit 		string		`json:"penerbit"`
	Pengarang 		string		`json:"pengarang"`
	JumlahHalaman 	int32		`json:"jumlah_halaman"`
	TahunTerbit 	int32		`json:"tahun_terbit"`
	Stock 			int32			`json:"stock"`
	KategoriID 		uuid.UUID	`json:"kategori_id"`
	CreatedAt   	time.Time  	`json:"created_at"`
	UpdatedAt   	*time.Time  	`json:"updated_at"`
	DeletedAt   	*time.Time 	`json:"deleted_at" sql:"index"`
}

func (b *Buku) Prepare()  {
	b.ID = uuid.NewV4()
	b.CreatedAt = time.Now()
}
