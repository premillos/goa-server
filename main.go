package main

import (
	"os"

	"com.goa/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()

	app.Description = "GOA 服务端API"
	app.Usage = ""
	app.Name = "com.goa"

	app.Commands = []*cli.Command{
		cmd.Start(),
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
