package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestInitRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	username := "tom"
	router := InitRouter()

	t.Run("index", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
	})
	t.Run("UserSave", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/user/"+username, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
	})
	t.Run("UserSaveQuery", func(t *testing.T) {
		age := "18"
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+age, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, "用户:"+username+",年龄:"+age+"已经保存", w.Body.String())
	})
	t.Run("UserSaveWithNotAge", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
	})
}
