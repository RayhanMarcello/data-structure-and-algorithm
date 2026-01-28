package main

import (
	"belajaraja/database"
	"belajaraja/handler"
	"belajaraja/repository"
	"belajaraja/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Database()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := gin.Default()
	r.GET("/user", handler.GetAllUser)
	r.GET("/user/:id", handler.GetUserById)

	r.POST("/user", handler.CreateUser)

	r.POST("/soal", handler.CreateSoal)

	r.Run(":8080")
}
