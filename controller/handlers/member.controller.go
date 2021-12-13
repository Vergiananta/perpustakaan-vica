package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/controller"
	"perpustakaan/models"
	"perpustakaan/models/dto"
	"perpustakaan/usecase"
	"perpustakaan/utils/response"
)

type MemberController struct {
	router 		*gin.Engine
	response 	response.IResponder
	service 	usecase.IMemberUsecase
}

func (m *MemberController) InitRoutes() {
	u := m.router.Group("/member")
	u.POST("", m.CreateMember)
	u.POST("/auth/login", m.LoginMember)
}

func NewMemberController(router *gin.Engine, responder response.IResponder, services usecase.IMemberUsecase ) controller.IDelivery {
	return &MemberController{
		router, responder, services,
	}
}

func (m *MemberController) CreateMember(c *gin.Context)  {

	var member models.Member

	if err := c.ShouldBindJSON(&member); err != nil {
		m.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
		return
	}

	memberCreated, err := m.service.CreateMember(&member)

	if err != nil {
		m.response.SetContext(c).ErrorResponder(http.StatusBadRequest, "",err.Error())
		return
	}

	m.response.SetContext(c).SingleResponder(http.StatusCreated, memberCreated)
}

func (m *MemberController) LoginMember(c *gin.Context)  {

	var login dto.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		m.response.SetContext(c).ErrorResponder(http.StatusBadRequest,"",err.Error())
		return
	}

	memberLogin, err:= m.service.Login(&login)
	if err != nil {
		m.response.SetContext(c).ErrorResponder(http.StatusBadRequest,"",err.Error())
		return
	}

	m.response.SetContext(c).SingleResponder(http.StatusOK, memberLogin)
}

func (m MemberController) InactiveAccount(c *gin.Context)  {

}