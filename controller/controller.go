package controller

import (
	"gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
// POST
func CreateATOdo (ctx *gin.Context) {
	// 前端点击提交， 请求发到这里

	// 1. 从请求中把数据拿出来
	var todo models.Todo
	ctx.BindJSON(&todo)
	// 2. 存入数据库
	if err := models.CreateATOdo(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
	// 3. 返回响应

}

func GetTodoList(ctx *gin.Context) {
	// 查询todo
	if todoList,err := models.GetTodoList(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	} else {
			ctx.JSON(http.StatusOK, todoList)
		}
}

func UpdateATodo(ctx *gin.Context) {
	// 修改事项的ID
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : "id no exists",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : "id no exists",
		})
		return
	}
	if err := models.DeleteATode(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : "id deleted",
		})
	}
}