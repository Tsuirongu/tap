package service

import (
	"fmt"
	"github.com/Tsuirongu/tap/service/impl"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default()
	r.GET("/", impl.SayHello)
	r.POST("/user/login", impl.Login)
	r.POST("/task/list", impl.TaskList)
	r.POST("/task/add", impl.TaskAdd)
	if err := r.Run(); err != nil {
		panic(fmt.Sprintf("start server fail: %v", err))
	} // listen and serve on 0.0.0.0:8080
}
