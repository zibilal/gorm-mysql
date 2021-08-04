package api

import (
	"github.com/go-playground/assert/v2"
	"gorm-mysql/appctx"
	"gorm-mysql/engine/apiengine"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test404NotFound(t *testing.T) {

	a := &Api{
		&appctx.AppCtx {
			ApiEngine: apiengine.InitApiEngine(),
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	router := a.Context.ApiEngine.Engine
	router.ServeHTTP(w, req)

	expected := `{"code":"NOT_FOUND","message":"Endpoint /ping Not Found"}`


	assert.Equal(t, 404, w.Code)
	assert.Equal(t, expected, w.Body.String())
}