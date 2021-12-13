package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/controller"
	"perpustakaan/models"
	"perpustakaan/usecase"
	"perpustakaan/utils/response"
)

type BukuController struct {
	router         *gin.Engine
	response 		response.IResponder
	service 	   usecase.IBukuUsecase
}

func (b *BukuController) InitRoutes() {
	u := b.router.Group("/buku")
	u.POST("", b.CreateNewBook)
}

func NewBookController(router *gin.Engine, responder response.IResponder,service usecase.IBukuUsecase) controller.IDelivery {
	return &BukuController{
		router, responder,service,
	}
}

func (b *BukuController) CreateNewBook(c *gin.Context) {

	var buku models.Buku

	if err := c.ShouldBindJSON(&buku); err != nil {
		b.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	bookCreated, err := b.service.CreateBook(&buku)

	if err != nil {
		b.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	b.response.SetContext(c).SingleResponder(http.StatusCreated, bookCreated)
}