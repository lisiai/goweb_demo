package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan.com/web_demo/log"
)

var R *gin.Engine

func Init() {
	SetUp()
	RouterG1()
}

func SetUp() {
	R = gin.New()
	R.Use(log.GinLogger(), log.GinRecovery(true))
}

func RouterG1() {
	R.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "成功"})
	})

}

func RouterG2() {

}
