package test

import (
	"GinHello/init"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := init.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	username := "lisi"
	router := init.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}
