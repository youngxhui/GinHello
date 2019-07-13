package test

import (
	"GinHello/initRouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// router("/") get 测试
func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())

}

// router("/") post 测试
func TestIndexPostRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin post method", w.Body.String())
}

// router("/") put 测试
func TestIndexPutRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodPut, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin put method", w.Body.String())
}

// router("/") delete 测试
func TestIndexDeleteRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin delete method", w.Body.String())
}

// router("/") patch 测试
func TestIndexPatchRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodPatch, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin patch method", w.Body.String())
}

// router("/") head 测试
func TestIndexHeadRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodHead, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin head method", w.Body.String())
}

// router("/") options 测试
func TestIndexOptionsRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodOptions, "/", nil)
	router.ServeHTTP(w, getReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin options method", w.Body.String())
}
