package router

import (
	"github.com/gin-gonic/gin"
	"perpustakaan/connect"
	"perpustakaan/controller/handlers"
)

type perpus struct {
	connect connect.Connect
	router  *gin.Engine
}

func (p *perpus) Run() {
	handlers.NewAppRouter(p.router, p.connect).InitMainRoutes()
	if err := p.router.Run(p.connect.ApiServer()); err != nil {
		panic(err)
	}
}

func NewPerpusApp() *perpus {
	r := gin.Default()
	var appConnect = connect.NewConnect()
	return &perpus{connect: appConnect, router: r}
}
