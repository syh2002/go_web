package rounters

import (
	"gin/controller"
	"github.com/gin-gonic/gin"
)

func SetupRounter() *gin.Engine {
	r := gin.Default()
	// 静态文件位置
	r.Static("/static", "static")
	// html文件位置
	r.LoadHTMLGlob("templates/*")
	// 首页
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 代办事项

		// 添加
		v1Group.POST("/todo", controller.CreateATOdo)
		// 查看 所有
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("todo/:id", controller.DeleteATodo)
	}
	return r
}