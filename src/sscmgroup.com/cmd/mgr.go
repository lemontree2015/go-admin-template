package cmd

import (
	"context"
	"github.com/urfave/cli/v2"
	"sscmgroup.com/app"
)

var mgrVersion = "1.0.0.1"

func NewMgrCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "mgr",
		Usage: "运行 mgr web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "casbin的访问控制模型(.conf)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.RunMgr(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				app.SetVersion(mgrVersion))
		},
	}
}
