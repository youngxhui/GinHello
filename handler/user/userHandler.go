package user

import (
	"GinHello/config"
	"GinHello/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateJwt(ctx *gin.Context) {
	// 获取用户
	user := &model.User{}
	result := &model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}
	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	u := user.QueryByUsername()
	if u.Password == user.Password {
		expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
		claims := jwt.StandardClaims{
			Audience:  user.Username,     // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        string(user.ID),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin hello",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		var jwtSecret = []byte(config.Secret)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
			result.Message = "登录成功"
			result.Data = "Bearer " + token
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		} else {
			result.Message = "登录失败"
			result.Code = http.StatusOK
			ctx.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusOK
		ctx.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}

func Register(ctx *gin.Context) {
	user := model.User{}
	result := &model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}
	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}

	if user.Insert() {
		result.Message = "注册成功"
		result.Code = http.StatusOK
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	} else {
		result.Message = "注册失败"
		result.Code = http.StatusOK
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}
