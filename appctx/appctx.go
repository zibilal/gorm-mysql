package appctx

import (
	"gorm-mysql/engine/apiengine"
	db2 "gorm-mysql/engine/dbengine"
)

var AppContext *AppCtx

type AppCtx struct {
	Host string
	DbContext *db2.DbEngine
	ApiEngine *apiengine.ApiEngine
	User map[string]string
}

func InitAppContext()  {
	AppContext = new(AppCtx)
}
