package usecase

import (
	"perpustakaan/models"
	"perpustakaan/models/dto"
	"perpustakaan/repositories"
)

type IMemberUsecase interface {
	CreateMember(newMember *models.Member) (*models.Member, error)
	Login(login *dto.LoginRequest) (string,error)
}



type memberUsecase struct {
	memberRepo repositories.IMemberRepository
}

func (m *memberUsecase) Login(login *dto.LoginRequest) (string, error) {
	return m.memberRepo.Login(login)
}

func (m *memberUsecase) CreateMember(newMember *models.Member) (*models.Member, error) {
	newMember.Prepare()
	return m.memberRepo.CreateMember(newMember)
}

func NewMemberUsecase(repo repositories.IMemberRepository) IMemberUsecase  {
	return &memberUsecase{repo}
}