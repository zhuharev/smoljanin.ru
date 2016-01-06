package catalog

import (
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
)

func Index(c *middleware.Context) {
	c.HTML(200, "catalog/index")
}

func Submit(c *middleware.Context) {
	if c.Req.Method == "POST" {
		submitPost(c)
		return
	}
	c.HTML(200, "catalog/submit")
}

func submitPost(c *middleware.Context) {
	c.Flash.Success(c.Query("link"), true)
	c.HTML(200, "catalog/submit")
}
