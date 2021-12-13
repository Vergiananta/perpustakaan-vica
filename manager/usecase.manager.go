package manager

import (
	"perpustakaan/connect"
	"perpustakaan/usecase"
)

type UsecaseManager interface {
	BukuUsecase() usecase.IBukuUsecase
	MemberUsecase() usecase.IMemberUsecase
	ReturnUsecase() usecase.IPengembalianUsecase
	PeminjamanUsecase() usecase.IPeminjamanUsecase
	DetailPinjamUsecase() usecase.IDetailPinjam
	TransactionsUsecase() usecase.ITransactionUsecase
	KategoriUsecase() usecase.IKategoriUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) KategoriUsecase() usecase.IKategoriUsecase {
	return usecase.NewKategoriUsecase(u.repo.KategoriRepo())
}

func (u *usecaseManager) TransactionsUsecase() usecase.ITransactionUsecase {
	return usecase.NewTransactionUsecase(u.PeminjamanUsecase(),u.ReturnUsecase(),u.DetailPinjamUsecase(), u.BukuUsecase())
}

func (u *usecaseManager) ReturnUsecase() usecase.IPengembalianUsecase {
	return usecase.NewPengembalianUsecase(u.repo.ReturnRepo())
}

func (u *usecaseManager) PeminjamanUsecase() usecase.IPeminjamanUsecase {
	return usecase.NewPeminjamanUsecase(u.repo.PeminjamanRepo())
}

func (u *usecaseManager) DetailPinjamUsecase() usecase.IDetailPinjam {
	return usecase.NewDetailPinjam(u.repo.DetailPinjamRepo())
}

func (u *usecaseManager) MemberUsecase() usecase.IMemberUsecase {
	return usecase.NewMemberUsecase(u.repo.MemberRepo())
}

func (u *usecaseManager) BukuUsecase() usecase.IBukuUsecase {
	return usecase.NewBukuUsecase(u.repo.BukuRepo())
}

func NewUsecaseManager(connect connect.Connect) UsecaseManager {
	return &usecaseManager{ NewRepoManager(connect)}
}
