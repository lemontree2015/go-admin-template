package cmd

import (
	"context"
	"github.com/urfave/cli/v2"
	"sscmgroup.com/task"
	"sscmgroup.com/task/tasks"
)

var tks []*cli.Command
var cleanFn func()

func init() {
	tks = append(tks, tasks.UserTasks, tasks.BookTasks)
}

func NewTaskCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "运行task服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件",
				Required: true,
			},
		},
		Before: func(c *cli.Context) (err error) {
			//注册配制文件
			cleanFn, err = task.Init(ctx,
				task.SetConfigFile(c.String("conf")),
			)
			return nil
		},
		After: func(c *cli.Context) error {
			cleanFn()
			return nil
		},
		Subcommands: tks,
	}
}
