package repositories

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"perpustakaan/middlewares"
	"perpustakaan/models"
	"perpustakaan/models/dto"
)

const (
	CreateNewMember = `insert into member(id, name, address, email, password, created_at) values($1,$2,$3,$4,$5,$6)`
	FindByEmail     = `select * from member m where email=$1`
)

type IMemberRepository interface {
	CreateMember(newMember *models.Member) (*models.Member, error)
	Login(login *dto.LoginRequest) (string, error)
}

type memberRepo struct {
	sql  *sql.DB
	stmt map[string]*sql.Stmt
}

func (m *memberRepo) Login(login *dto.LoginRequest) (string, error) {
	var err error
	rows, errs := m.stmt["LoginStatement"].Query(login.Email)
	if errs != nil {
		return "", errs
	}

	var account models.Member
	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Name, &account.Address, &account.Email, &account.Password, &account.CreatedAt, &account.UpdatedAt, &account.DeletedAt)
		if err != nil || account.Name == "" {
			return "", err
		}

		err = models.VerifyPassword(account.Password, login.Password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return "", err
		}
	}
	return middlewares.CreateToken(&account)
}

func (m *memberRepo) CreateMember(newMember *models.Member) (*models.Member, error) {

	_, err := m.sql.Exec(CreateNewMember, newMember.ID, newMember.Name, newMember.Address, newMember.Email, newMember.Password, newMember.CreatedAt)
	if err != nil {
		return nil, err
	}

	return newMember, nil
}

func NewMemberRepository(db *sql.DB) IMemberRepository {
	stmt := map[string]*sql.Stmt{}
	queries, err := db.Prepare(FindByEmail)
	if err != nil {
		panic(err)
	}
	stmt["LoginStatement"] = queries
	return &memberRepo{
		db,
		stmt,
	}
}
