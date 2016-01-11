package controllers

import (
	"fmt"
	"github.com/zhuharev/smoljanin.ru/modules/middleware"
	"github.com/zhuharev/smoljanin.ru/modules/uuid"
	"io"
	"os"
	"path/filepath"
)

func Home(c *middleware.Context) {
	c.HTML(200, "home")
}

func Upload(c *middleware.Context) {
	c.Req.ParseMultipartForm(32 << 20)
	file, _, err := c.Req.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	name := uuid.NewV4().String()

	//.fmt.Fprintf(w, "%v", handler.Header)
	err = os.MkdirAll("data/uploads/"+makePaths(name), 0777)
	if err != nil {
		return
	}
	f, err := os.OpenFile("data/uploads/"+makePaths(name)+"/"+name, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	c.JSON(200, makePaths(name))
}

func makePaths(p string) string {
	return filepath.Join(string(p[0]), string(p[1]), string(p[2]))
}
