package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}
