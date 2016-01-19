package catalog

import (
	"github.com/Unknwon/paginater"
	"github.com/fatih/color"
	"github.com/zhuharev/smoljanin.ru/models"
	"github.com/zhuharev/smoljanin.ru/modules/linkpreview"
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
	"time"
)

func Feed(c *middleware.Context) {
	sid := c.ParamsInt64(":id")
	p := c.ParamsInt(":p")
	if p < 1 {
		p = 1
	}
	feed, e := models.GetSiteFeed(sid, p)
	if e != nil {
		color.Red("%s", e)
	}

	cnt, e := models.SiteFeedCount(sid)
	if e != nil {
		color.Red("%s", e)
	}

	go models.SiteFetchNewFeed(sid)
	c.Data["feed"] = feed
	c.Data["paginater"] = paginater.New(int(cnt), 10, p, 5)
	c.Data["SiteId"] = sid
	c.HTML(200, "catalog/feed")
}

func SetFeed(c *middleware.Context) {
	id := c.ParamsInt64(":id")
	val := c.Query("hf")
	hf := false
	if val == "1" {
		hf = true
	}
	e := models.SiteSetHasFeed(id, hf)
	if e != nil {
		c.Flash.Error(e.Error())
	}
	c.Redirect(c.URLFor("cat_item", ":id", c.Params(":id")))
}

func FeedShow(c *middleware.Context) {
	sf, e := models.GetSiteFeedItem(c.ParamsInt64(":feedId"))
	if e != nil {
		color.Red("%s", e)
	}
	a, e := linkpreview.Articler.ParseArticle(sf.Source, []byte(sf.Body))
	if e != nil {
		color.Red("%s", e)
	}
	if !sf.Published.Equal(a.Published.In(time.Now().Location())) {
		sf.Published = a.Published.In(time.Now().Location())
		models.SaveSiteFeed(sf)
	}
	sf.Body = a.Text
	c.Data["feed"] = sf
	c.HTML(200, "catalog/feedshow")
}
