package manager

import (
	"database/sql"
	"perpustakaan/connect"
	"perpustakaan/repositories"
)

type RepoManager interface {
	BukuRepo() repositories.IBukuRepository
	MemberRepo() repositories.IMemberRepository
	ReturnRepo() repositories.IPengembalianRepo
	PeminjamanRepo() repositories.IPeminjamanRepo
	DetailPinjamRepo() repositories.IDetailPinjamRepo
	KategoriRepo() repositories.IKategoriRepo
}

type repoManager struct {
	sqlConn *sql.DB
	connect connect.Connect
}

func (r *repoManager) KategoriRepo() repositories.IKategoriRepo {
	return repositories.NewKategoriRepo(r.sqlConn)
}

func (r *repoManager) ReturnRepo() repositories.IPengembalianRepo {
	return repositories.NewPengembalianRepo(r.sqlConn)
}

func (r *repoManager) PeminjamanRepo() repositories.IPeminjamanRepo {
	return repositories.NewPeminjamanRepo(r.sqlConn)
}

func (r *repoManager) DetailPinjamRepo() repositories.IDetailPinjamRepo {
	return repositories.NewDetailPinjamRepo(r.sqlConn)
}

func (r *repoManager) MemberRepo() repositories.IMemberRepository {
	return repositories.NewMemberRepository(r.sqlConn)
}

func (r *repoManager) BukuRepo() repositories.IBukuRepository {
	return repositories.NewBukuRepository(r.sqlConn)
}

func NewRepoManager(connect connect.Connect) RepoManager{
	return &repoManager{connect.SqlDb(), connect}
}
