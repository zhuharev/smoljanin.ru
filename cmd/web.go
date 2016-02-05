package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/session"
	"github.com/zhuharev/smoljanin.ru/modules/base"
	"github.com/zhuharev/smoljanin.ru/modules/since"
	"gopkg.in/macaron.v1"
	"html/template"
	"time"

	"github.com/zhuharev/smoljanin.ru/controllers"
	"github.com/zhuharev/smoljanin.ru/controllers/catalog"
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
	"github.com/zhuharev/smoljanin.ru/modules/setting"
)

var CmdWeb = cli.Command{
	Name:  "web",
	Usage: "Start Gogs web server",
	Description: `Gogs web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		stringFlag("port, p", "3000", "Temporary port number to prevent conflict"),
		stringFlag("config, c", "custom/conf/app.ini", "Custom configuration file path"),
	},
}

func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Renderer(macaron.RenderOptions{Layout: "layout",
		Funcs: []template.FuncMap{{
			"markdown": base.Markdown,
			"raw":      func(s string) template.HTML { return template.HTML(s) },
			"momentDiff": func(t time.Time) string {
				return since.Since(t)
			},
		}}}))
	/*	m.Use(func(c *macaron.Context) {
		if strings.HasSuffix(c.Req.URL.Path, ".json") {
			color.Green("JSON")

			c.Req.Request.URL

			c.Req.URL.Path = strings.TrimSuffix(c.Req.URL.Path, ".json")
			c.Req.URL.RawPath = strings.TrimSuffix(c.Req.URL.RawPath, ".json")
			c.Req.RequestURI = c.Req.URL.RequestURI()

			c.Data["json"] = true
		}
		c.Next()
	})*/
	m.Use(cache.Cacher())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	m.Use(macaron.Static("static"))
	m.Use(macaron.Static("data/uploads"))
	m.Use(macaron.Static("data/public", macaron.StaticOptions{Prefix: "public"}))

	m.Use(i18n.I18n(i18n.Options{
		Langs: []string{"en-US", "ru-RU"},
		Names: []string{"English", "Русский"},
	}))

	m.Use(middleware.Contexter())

	return m
}

func runWeb(c *cli.Context) {
	if c.IsSet("config") {
		setting.CustomConf = c.String("config")
	}
	controllers.GlobalInit()
	m := newMacaron()

	m.Get("/", controllers.Home)
	m.Get("/feed.json", catalog.AllFeedJSON)
	m.Get("/ok.json", func(c *macaron.Context) { c.JSON(200, "ok") })

	m.Group("/cat", func() {
		m.Get("/", catalog.Index)
		m.Get("/page/:p", catalog.Index)
		m.Get("/:id", catalog.Show).Name("cat_item")
		m.Get("/:id/feed", catalog.Feed)
		m.Get("/:id/feed/page/:p", catalog.Feed)
		m.Get("/:id/feed/:feedId", catalog.FeedShow)
		m.Get("/screen/:id", catalog.Screen)

		m.Get("/:id/setfeed", catalog.SetFeed).Name("set_feed")

		m.Any("/submit", catalog.Submit)
	})

	m.Get("/resize/*", catalog.Resize)

	m.Post("/upload", controllers.Upload)

	m.Run(5000)
}
