// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"github.com/Unknwon/com"

	"gopkg.in/ini.v1"

	"github.com/zhuharev/smoljanin.ru/modules/log"
)

var (
	AttachmentMaxSize int64
	AppSubUrl         string
	GravatarSource    = "//en.gravatar.com"

	Cfg        *ini.File
	CustomConf string
)

func NewContext() {
	var (
		e error
	)
	Cfg, e = ini.Load("conf/app.ini")
	if e != nil {
		log.Fatal(4, "Fail to load conf %v", e)
	}
	if com.IsFile(CustomConf) {
		if e = Cfg.Append(CustomConf); e != nil {
			log.Fatal(4, "Fail to load custom conf '%s': %v", CustomConf, e)
		}
	} else {
		log.Warn("Custom config (%s) not found, ignore this if you're running first time", CustomConf)
	}
	Cfg.NameMapper = ini.AllCapsUnderscore
}
