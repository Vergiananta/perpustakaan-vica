package handlers

import (
	"net/http"
	"perpustakaan/controller"
	"perpustakaan/models"
	"perpustakaan/usecase"
	"perpustakaan/utils/response"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type KategoriController struct {
	router   *gin.Engine
	response response.IResponder
	service  usecase.IKategoriUsecase
}

func (k *KategoriController) InitRoutes() {
	u := k.router.Group("")
	u.POST("", k.CreateKategori)
}

func NewKategoriController(router *gin.Engine, responder response.IResponder, service usecase.IKategoriUsecase) controller.IDelivery {
	return &KategoriController{
		router, responder, service,
	}
}

func (k *KategoriController) CreateKategori(c *gin.Context) {
	var newCategory models.Kategori

	categories := []models.Kategori{
		{
			ID:        uuid.NewV4(),
			Name:      "Novel",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.NewV4(),
			Name:      "Magazine",
			CreatedAt: time.Now(),
		},
	}

	for i, _ := range categories {
		newCategory = categories[i]
	}
	kategoriPerpus, err := k.service.CreateKategori(&newCategory)
	if err != nil {
		k.response.ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	k.response.SetContext(c).SingleResponder(http.StatusCreated, kategoriPerpus)
}
