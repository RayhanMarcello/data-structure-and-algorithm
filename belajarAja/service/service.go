package service

import (
	"belajaraja/entity"
	"belajaraja/repository"
	"fmt"
)

type Service interface {
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)

	CreateUser(username string, email string, age int) (entity.User, error)
	CreateSoal(soal string, jawaban string) (entity.Ujian, error)
}

type service struct {
	repo repository.Repository
}

func NewService(srv repository.Repository) *service {
	return &service{srv}
}

func (s *service) GetAllUser() ([]entity.User, error) {
	return s.repo.GetAllUser()
}

func (s *service) GetUserById(id int) (entity.User, error) {
	if id <= 0 {
		return entity.User{}, fmt.Errorf("invalid id!")
	}
	return s.repo.GetUserById(id)
}

func (s *service) CreateUser(username string, email string, age int) (entity.User, error) {
	if email == "" {
		return entity.User{}, fmt.Errorf("")
	}
	if username == "" {
		return entity.User{}, fmt.Errorf("")
	}

	if age < 0 {
		return entity.User{}, fmt.Errorf("")
	}

	users := entity.User{
		Username: username,
		Email:    email,
		Age:      age,
	}

	return s.repo.CreateUser(users)
}

func (s *service) CreateSoal(soal string, jawaban string) (entity.Ujian, error) {
	if soal == "" {
		return entity.Ujian{}, fmt.Errorf("masukan soal!")
	}
	if jawaban == "" {
		return entity.Ujian{}, fmt.Errorf("masukan jawaban !")
	}

	ujian := entity.Ujian{
		Soal:    soal,
		Jawaban: jawaban,
	}
	res, err := s.repo.CreateSoal(ujian)
	if err != nil {
		fmt.Println(err)
	}
	return res, err

}
