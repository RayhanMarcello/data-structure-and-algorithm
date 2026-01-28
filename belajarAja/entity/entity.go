package entity

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Email    string
	Age      int
}
type Ujian struct {
	Soal    string
	Jawaban string
}
