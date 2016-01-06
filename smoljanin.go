package main

import (
	"github.com/codegangsta/cli"
	"github.com/zhuharev/smoljanin.ru/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Smoljanin.ru website"
	//app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
