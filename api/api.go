package api

import (
	"gorm-mysql/appctx"
	"log"
	"net/http"
)

type ControlRouter interface {
	Router(string, string)
}

type Api struct {
	Context *appctx.AppCtx
}

func NewApi(ctx *appctx.AppCtx) *Api {
	a := &Api{Context: ctx}
	return a
}

func (c *Api) RegisterRouter(groupName, version string, rc ControlRouter) {
	rc.Router(groupName, version)
}

func (c *Api) Run() {
	if err:=http.ListenAndServe(
		c.Context.Host,
		c.Context.ApiEngine.Engine); err != nil {
		log.Fatal(err)
	}
}
