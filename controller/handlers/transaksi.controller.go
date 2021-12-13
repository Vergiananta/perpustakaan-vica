package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/controller"
	"perpustakaan/middlewares"
	"perpustakaan/models/dto"
	"perpustakaan/usecase"
	"perpustakaan/utils/response"
)

type TransactionController struct {
	router   *gin.Engine
	response response.IResponder
	service  usecase.ITransactionUsecase
}

func (t *TransactionController) InitRoutes() {
	u := t.router.Group("/transactions")
	u.POST("/borrow", middlewares.SetMiddlewareAuthentication(t.ReturnBook))
}

func NewTransactionController(router *gin.Engine, responder response.IResponder, services usecase.ITransactionUsecase) controller.IDelivery {
	return &TransactionController{router, responder, services}
}

func (t *TransactionController) ReturnBook(c *gin.Context) {

	var trx dto.PinjamRequest

	if err := c.ShouldBindJSON(&trx); err != nil {
		t.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}

	borrow, err := t.service.PinjamBuku(&trx)

	if err != nil {
		t.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}
	t.response.SetContext(c).SingleResponder(http.StatusCreated, borrow)
}
