package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nihcet/zlog/src/handler"
	handlerwrapper_util "github.com/nihcet/zlog/util/handler-wrapper"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	h := handler.InitialHandler()

	r.GET("/ping", handlerwrapper_util.Get(h.Ping))
	r.GET("/pong", handlerwrapper_util.GetOne(h.Pong))

	return r
}
