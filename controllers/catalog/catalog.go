package catalog

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/Unknwon/paginater"
	"github.com/disintegration/imaging"
	"github.com/fatih/color"
	"github.com/nfnt/resize"
	"github.com/zhuharev/smoljanin.ru/models"
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
	"github.com/zhuharev/smoljanin.ru/modules/shot"
	"github.com/zhuharev/smoljanin.ru/modules/uuid"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	c.Data["Title"] = site.Title
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

func Screen(c *middleware.Context) {
	id := c.ParamsInt64(":id")
	name, e := screen(id)
	if e != nil {
		color.Red("%s", e)
	}
	c.Redirect("/" + makePaths(name) + "/" + name)
}

func screen(siteId int64) (string, error) {
	site, e := models.GetSite(siteId)
	if e != nil {
		return "", e
	}

	rdr, e := shot.Shot(site.Domain)

	name := uuid.NewV4().String()

	//.fmt.Fprintf(w, "%v", handler.Header)
	e = os.MkdirAll("data/uploads/"+makePaths(name), 0777)
	if e != nil {
		return "", e
	}
	f, e := os.OpenFile("data/uploads/"+makePaths(name)+"/"+name, os.O_WRONLY|os.O_CREATE, 0777)
	if e != nil {
		return "", e
	}
	defer f.Close()
	io.Copy(f, rdr)

	fp := filepath.Join("/", makePaths(name), "/", name)
	site.PreviewUrl = fp
	e = models.SaveSite(site)
	if e != nil {
		return "", e
	}
	return fp, nil
}

func Resize(c *middleware.Context) {
	w := c.QueryInt("width")
	p := c.Req.URL.Path
	color.Green("[path] %s", p)
	target := strings.TrimPrefix(p, "/resize")
	if strings.HasPrefix(target, "/cat/screen/") {
		var id int64
		fmt.Sscanf(target, "/cat/screen/%d?width=", &id)
		name, e := screen(id)
		if e != nil {
			color.Red("%s", e)
			return
		}
		target = name
	}
	savePath := "data/public/resize/" + com.ToStr(w) + "/" + target + ".jpg"
	if com.IsFile(savePath) {
		f, e := os.Open(savePath)
		if e != nil {
			color.Red("%s", e)
		}
		c.Resp.Header().Set("Content-Type", "image/jpeg")
		io.Copy(c.Resp, f)
		return
	}

	img, e := imaging.Open(filepath.Join("data/uploads", target))
	if e != nil {
		color.Red("%s", e)
		return
	}
	i := resize.Resize(uint(w), uint(w/2), img, resize.Bilinear)
	dir := filepath.Dir(target)
	os.MkdirAll(filepath.Join("data/public/resize/"+com.ToStr(w)+"/", dir), 0777)
	e = imaging.Save(i, savePath)
	if e != nil {
		color.Red("%s", e)
	}
	c.Resp.Header().Set("Content-Type", "image/jpeg")
	e = jpeg.Encode(c.Resp, i, nil)
	if e != nil {
		color.Red("%s", e)
	}
	//c.Redirect("/public/resize/" + com.ToStr(w) + "/" + target + ".jpg")
}

func makePaths(p string) string {
	return filepath.Join(string(p[0]), string(p[1]), string(p[2]))
}
