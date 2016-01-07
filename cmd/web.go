package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"

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
	m.Use(macaron.Renderer(macaron.RenderOptions{Layout: "layout"}))
	m.Use(cache.Cacher())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	m.Use(macaron.Static("static"))
	m.Use(macaron.Static("data/uploads"))

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
	m.Get("/cat", catalog.Index)
	m.Get("/cat/page/:p", catalog.Index)
	m.Any("/cat/submit", catalog.Submit)
	m.Get("/cat/:id", catalog.Show)

	m.Run()
}
