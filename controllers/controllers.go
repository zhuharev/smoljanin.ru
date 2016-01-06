package controllers

import (
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
)

func Home(c *middleware.Context) {
	c.HTML(200, "home")
}
