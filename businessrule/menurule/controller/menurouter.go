package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-mysql/appctx"
)

type MenuVo struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type MenuRouter struct {
	addMenu func(input interface{}, output interface{}) error
	ctx *appctx.AppCtx
}

func NewMenuRouter(ctx *appctx.AppCtx) *MenuRouter {
	r := new(MenuRouter)
	r.ctx = ctx

	return r
}

func (r *MenuRouter) Router(groupName, version string) {
	menu := r.ctx.ApiEngine.Engine.Group(fmt.Sprintf("/%s/%s", groupName, version))
	menu.POST("/", func(c *gin.Context) {
		vo := struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Description string `json:"description"`
		}{}
		if err := c.ShouldBind(&vo); err != nil{
			c.JSON(400, gin.H{
				"code": "UNKNOWN_DATA_FORMAT",
				"message": fmt.Sprintf("Unknown Data Format: %s", err.Error()),
			})
			return
		}
	})
}
