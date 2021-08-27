package main

import (
	"fmt"
	"go.uber.org/zap"
	"turan.com/web_demo/dao/mysql"
	"turan.com/web_demo/dao/redis"
	"turan.com/web_demo/log"
	"turan.com/web_demo/router"
	set "turan.com/web_demo/settings"
)

func main() {
	//1.加载配置信息
	err := set.Init()
	if err != nil {
		panic(fmt.Errorf("Failed to load configuration err for", err))
	}
	//2,初始化日志:需增加分割功能
	err = log.Init()
	if err != nil {
		panic(fmt.Errorf("Mysql initialization failed. Procedure,err "))
	}
	defer zap.L().Sync()
	//3，初始化mysql
	err = mysql.Init()
	if err != nil {
		panic(fmt.Errorf("Mysql initialization failed. Procedure,err "))
	}
	defer mysql.Close()
	//4，初始化redis
	err = redis.Init()
	defer redis.Close()
	if err != nil {
		panic(fmt.Errorf("redis initialization failed ,procedure,err"))
	}
	//5，注册路由
	router.Init()
	//6，启动服务
	router.Run()

}
