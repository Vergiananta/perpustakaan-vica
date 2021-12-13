package usecase

import (
	"errors"
	"perpustakaan/models"
	"perpustakaan/models/dto"
	"time"
)

type ITransactionUsecase interface {
	PinjamBuku(pinjam *dto.PinjamRequest) (*dto.PinjamResponse, error)
	ReturnBuku(returnBook *dto.ReturnRequest) (*dto.ReturnResponse, error)
}

type transactionUsecase struct {
	servicePinjam       IPeminjamanUsecase
	servicePengembalian IPengembalianUsecase
	serviceDetailPinjam IDetailPinjam
	serviceBuku         IBukuUsecase
}

func (t *transactionUsecase) ReturnBuku(returnBook *dto.ReturnRequest) (*dto.ReturnResponse, error) {
	today := time.Now()
	panic(today)
}

func (t *transactionUsecase) PinjamBuku(pinjam *dto.PinjamRequest) (*dto.PinjamResponse, error) {
	today := time.Now()

	pinjamBuku := models.Peminjaman{
		MemberID:      pinjam.MemberID,
		TanggalPinjam: today.Format("2006/01/02"),
		LamaPinjam:    pinjam.LamaPinjam,
		Keterangan:    pinjam.Keterangan,
		Status:        models.DIPINJAM,
	}
	pinjamBuku.Prepare()
	_, err := t.servicePinjam.CreatePinjam(&pinjamBuku)
	if err != nil {
		return nil, err
	}

	detailPinjam := models.PeminjamanDetail{
		PeminjamanID: pinjamBuku.ID,
		BukuID:       pinjam.BukuID,
		Quantity:     pinjam.Quantity,
	}
	detailPinjam.Prepare()
	t.serviceDetailPinjam.CreateDetailPinjam(&detailPinjam)
	//end:= today.AddDate(0,0,pinjam.LamaPinjam)

	books, err := t.serviceBuku.FindBookById(pinjam.BukuID)
	books.Stock -= pinjam.Quantity

	returnBook := models.Pengembalian{
		PeminjamanID:   pinjamBuku.ID,
		TanggalKembali: today.Format("2006/01/02"),
	}
	returnBook.Prepare()
	returnB, errR := t.servicePengembalian.CreatePengembalian(&returnBook)
	if errR != nil {
		return nil, errR
	}

	if books.Stock >= 0 {
		_, errU := t.serviceBuku.UpdateStock(books)
		if errU != nil {
			return nil, errU
		}
		return &dto.PinjamResponse{
			ID:             pinjamBuku.ID,
			TanggalPinjam:  today.Format("2006/01/02"),
			TanggalKembali: returnB.TanggalKembali,
			BukuID:         books.ID,
			Quantity:       detailPinjam.Quantity,
		}, nil
	} else {
		return nil, errors.New("Book out of stock")
	}
}

func NewTransactionUsecase(servPinjam IPeminjamanUsecase, servKembali IPengembalianUsecase, servDetail IDetailPinjam, servBook IBukuUsecase) ITransactionUsecase {
	return &transactionUsecase{servPinjam, servKembali, servDetail, servBook}
}
