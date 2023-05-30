package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}
