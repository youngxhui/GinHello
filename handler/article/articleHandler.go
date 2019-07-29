package article

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 提交新的文章内容
// @Id 1
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 添加失败
// @Router /article [post]
func Insert(context *gin.Context) {
	article := model.Article{}
	var id = -1
	var message = "数据添加失败"
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
		message = "数据添加成功"
	}
	result := model.Result{
		Code:    http.StatusOK,
		Message: message,
		Data: gin.H{
			"id": id,
		},
	}
	context.JSON(http.StatusOK, result)
}

// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 {object} model.Result 成功后返回值
// @Router /article/{id} [get]
func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.Article{
		Id: i,
	}

	art := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"article": art,
	})

}

// @Summary 获取所有的文章
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.Result 成功后返回值
// @Router /articles [get]
func GetAll(context *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

// @Summary 通过id删除指定文章类型
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200  object model.Result 成功后返回值
// @Router /article/{id} [delete]
func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.Article{Id: i}
	article.DeleteOne()
}
