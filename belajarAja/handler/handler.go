package handler

import (
	"belajaraja/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	srv service.Service
}

func NewHandler(h service.Service) *handler {
	return &handler{h}
}

func (h *handler) GetAllUser(c *gin.Context) {
	user, err := h.srv.GetAllUser()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot reach users",
		})
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    user,
	})

}

func (h *handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	idParams, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
	}

	data, err := h.srv.GetUserById(idParams)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

type reqBody struct {
	Username string
	Email    string
	Age      int
}

func (h *handler) CreateUser(c *gin.Context) {
	var req reqBody
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "bad req",
			"data":    err,
		})
	}

	create, err := h.srv.CreateUser(req.Email, req.Username, req.Age)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "bad req",
			"data":    err,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": create,
	})
}

// soal section

type reqSoal struct {
	Soal    string
	Jawaban string
}

func (h *handler) CreateSoal(c *gin.Context) {
	var reqSoal reqSoal
	if err := c.ShouldBindBodyWithJSON(&reqSoal); err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	soal, err := h.srv.CreateSoal(reqSoal.Soal, reqSoal.Jawaban)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"datas":   soal,
	})
}
