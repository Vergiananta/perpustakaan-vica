package handlers

import (
	"github.com/gin-gonic/gin"
	"perpustakaan/connect"
	"perpustakaan/manager"
	"perpustakaan/utils/response"
)

type appRouter struct {
	router    *gin.Engine
	connect connect.Connect
}


func (r *appRouter) InitMainRoutes() {
	serviceManager := manager.NewUsecaseManager(r.connect)
	NewMemberController(r.router, response.NewJsonResponder(),serviceManager.MemberUsecase()).InitRoutes()
	NewKategoriController(r.router,response.NewJsonResponder(),serviceManager.KategoriUsecase()).InitRoutes()
	NewBookController(r.router, response.NewJsonResponder(),serviceManager.BukuUsecase()).InitRoutes()
	NewTransactionController(r.router,response.NewJsonResponder(),serviceManager.TransactionsUsecase()).InitRoutes()
}

func NewAppRouter(app *gin.Engine, connect connect.Connect) *appRouter {
	return &appRouter{
		app,
		connect,
	}
}