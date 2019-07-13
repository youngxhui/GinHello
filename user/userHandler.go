package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Save(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 通过 query 方法进行获取参数
func SaveByQuery(context *gin.Context) {
	username := context.Query("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}
