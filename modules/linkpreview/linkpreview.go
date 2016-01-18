package linkpreview

import (
	"github.com/sisteamnik/articler"
)

var (
	Articler *articler.Articler
)

func NewContext() {
	var (
		cnpath = "conf/articler.conf"
		e      error
	)

	cnf := new(articler.Config)
	cnf.DefaultArticleParserConf = cnpath

	Articler, e = articler.New(cnf)
	if e != nil {
		panic(e)
	}
}
