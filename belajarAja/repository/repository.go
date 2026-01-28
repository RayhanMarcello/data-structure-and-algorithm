package repository

import (
	"belajaraja/entity"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)

	CreateUser(user entity.User) (entity.User, error)

	CreateSoal(soal entity.Ujian) (entity.Ujian, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]entity.User, error) {
	var user []entity.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *repository) GetUserById(id int) (entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error
	if err != nil {
		fmt.Println("db err from repos")
	}

	return user, err
}

func (r *repository) CreateUser(user entity.User) (entity.User, error) {
	users := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
	err := r.db.Create(&users).Error
	if err != nil {
		fmt.Println("cannto create user from repository layer")
	}
	return users, err
}

func (r *repository) CreateSoal(ujian entity.Ujian) (entity.Ujian, error) {
	soal := entity.Ujian{
		Soal:    ujian.Soal,
		Jawaban: ujian.Jawaban,
	}
	err := r.db.Create(&soal).Error
	if err != nil {
		fmt.Println("cant accsess db to create soal")
	}
	return soal, err

}
