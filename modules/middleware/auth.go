// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
/*"fmt"
"net/url"

"github.com/go-macaron/csrf"
"gopkg.in/macaron.v1"

"github.com/zhuharev/smoljanin.ru/models"
"github.com/zhuharev/smoljanin.ru/modules/auth"
"github.com/zhuharev/smoljanin.ru/modules/base"
"github.com/zhuharev/smoljanin.ru/modules/log"
"github.com/zhuharev/smoljanin.ru/modules/setting"*/
)

type ToggleOptions struct {
	SignInRequire  bool
	SignOutRequire bool
	AdminRequire   bool
	DisableCsrf    bool
}

// AutoSignIn reads cookie and try to auto-login.
/*func AutoSignIn(ctx *Context) (bool, error) {
	if !models.HasEngine {
		return false, nil
	}

	uname := ctx.GetCookie(setting.CookieUserName)
	if len(uname) == 0 {
		return false, nil
	}

	isSucceed := false
	defer func() {
		if !isSucceed {
			log.Trace("auto-login cookie cleared: %s", uname)
			ctx.SetCookie(setting.CookieUserName, "", -1, setting.AppSubUrl)
			ctx.SetCookie(setting.CookieRememberName, "", -1, setting.AppSubUrl)
		}
	}()

	u, err := models.GetUserByName(uname)
	if err != nil {
		if !models.IsErrUserNotExist(err) {
			return false, fmt.Errorf("GetUserByName: %v", err)
		}
		return false, nil
	}

	if val, _ := ctx.GetSuperSecureCookie(
		base.EncodeMD5(u.Rands+u.Passwd), setting.CookieRememberName); val != u.Name {
		return false, nil
	}

	isSucceed = true
	ctx.Session.Set("uid", u.Id)
	ctx.Session.Set("uname", u.Name)
	return true, nil
}

func Toggle(options *ToggleOptions) macaron.Handler {
	return func(ctx *Context) {
		// Cannot view any page before installation.
		if !setting.InstallLock {
			ctx.Redirect(setting.AppSubUrl + "/install")
			return
		}

		// Checking non-logged users landing page.
		if !ctx.IsSigned && ctx.Req.RequestURI == "/" && setting.LandingPageUrl != setting.LANDING_PAGE_HOME {
			ctx.Redirect(setting.AppSubUrl + string(setting.LandingPageUrl))
			return
		}

		// Redirect to dashboard if user tries to visit any non-login page.
		if options.SignOutRequire && ctx.IsSigned && ctx.Req.RequestURI != "/" {
			ctx.Redirect(setting.AppSubUrl + "/")
			return
		}

		if !options.SignOutRequire && !options.DisableCsrf && ctx.Req.Method == "POST" && !auth.IsAPIPath(ctx.Req.URL.Path) {
			csrf.Validate(ctx.Context, ctx.csrf)
			if ctx.Written() {
				return
			}
		}

		if options.SignInRequire {
			if !ctx.IsSigned {
				// Restrict API calls with error message.
				if auth.IsAPIPath(ctx.Req.URL.Path) {
					ctx.APIError(403, "", "Only signed in user is allowed to call APIs.")
					return
				}

				ctx.SetCookie("redirect_to", url.QueryEscape(setting.AppSubUrl+ctx.Req.RequestURI), 0, setting.AppSubUrl)
				ctx.Redirect(setting.AppSubUrl + "/user/login")
				return
			} else if !ctx.User.IsActive && setting.Service.RegisterEmailConfirm {
				ctx.Data["Title"] = ctx.Tr("auth.active_your_account")
				ctx.HTML(200, "user/auth/activate")
				return
			}
		}

		// Try auto-signin when not signed in.
		if !options.SignOutRequire && !ctx.IsSigned && !auth.IsAPIPath(ctx.Req.URL.Path) {
			succeed, err := AutoSignIn(ctx)
			if err != nil {
				ctx.Handle(500, "AutoSignIn", err)
				return
			} else if succeed {
				log.Trace("Auto-login succeed: %s", ctx.Session.Get("uname"))
				ctx.Redirect(setting.AppSubUrl + ctx.Req.RequestURI)
				return
			}
		}

		if options.AdminRequire {
			if !ctx.User.IsAdmin {
				ctx.Error(403)
				return
			}
			ctx.Data["PageIsAdmin"] = true
		}
	}
}
*/
