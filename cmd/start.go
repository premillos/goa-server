package cmd

import (
	"context"

	"com.goa/internal"
	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {

	return &cli.Command{
		Name:  "start",
		Usage: "启动服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "配置文件路径",
			},
		},
		Action: func(c *cli.Context) error {
			config := c.String("config")

			internal.Bootstrap(context.Background(), config)
			return nil
		},
	}
}
