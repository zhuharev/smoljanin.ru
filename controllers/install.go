package controllers

import (
	"github.com/zhuharev/smoljanin.ru/models"
	"github.com/zhuharev/smoljanin.ru/modules/menu"
	"github.com/zhuharev/smoljanin.ru/modules/setting"
)

func GlobalInit() {
	setting.NewContext()
	e := models.NewEngine()
	if e != nil {
		panic(e)
	}
	menu.NewContext()
}
