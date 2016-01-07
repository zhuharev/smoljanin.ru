package catalog

import (
	"fmt"
	"github.com/Unknwon/paginater"
	"github.com/fatih/color"
	"github.com/zhuharev/smoljanin.ru/models"
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
)

func Index(c *middleware.Context) {
	var (
		p = c.ParamsInt(":p")
	)
	sites, e := models.SiteList(p)
	if e != nil {
		color.Red("%s", e)
	}

	cnt, e := models.SiteCount()
	if e != nil {
		color.Red("%s", e)
	}

	c.Data["sites"] = sites
	c.Data["sites_count"] = cnt
	c.Data["paginater"] = paginater.New(int(cnt), 10, p, 5)
	c.HTML(200, "catalog/index")
}

func Show(c *middleware.Context) {
	site, e := models.GetSite(c.ParamsInt64(":id"))
	if e != nil {
		color.Red("%s", e)
	}
	c.Data["site"] = site
	c.HTML(200, "catalog/show")
}

func Submit(c *middleware.Context) {
	if c.Req.Method == "POST" {
		submitPost(c)
		return
	}
	c.HTML(200, "catalog/submit")
}

func submitPost(c *middleware.Context) {

	s, e := models.NewSite(c.Query("link"), c.Query("title"))
	if e != nil {
		c.Flash.Error(e.Error(), true)
		if c.QueryInt("ajax") == 1 {
			c.JSON(200, e.Error())
			return
		}
		c.HTML(200, "catalog/submit")
		return
	}

	c.Flash.Success(fmt.Sprintf("%d  добавлен", s.Id), true)
	if c.QueryInt("ajax") == 1 {
		c.JSON(200, fmt.Sprintf("%d  добавлен", s.Id))
		return
	}
	c.HTML(200, "catalog/submit")
}
