package tasks

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//use age go run main.go task -c 234 user del liangpang 55 345346346 tttt
//use age go run main.go task -c 234 user add -n liangpang
var UserTasks = &cli.Command{
	Name:  "user",
	Usage: "运行task user 服务",
	Before: func(c *cli.Context) error {
		//注册配制文件
		fmt.Println("user Before")
		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name:  "add",
			Usage: "add a user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
					Aliases:  []string{"n"},
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("user add ", c.String("name"))
				return nil
			},
		},
		{
			Name:  "del",
			Usage: "del a user",
			Action: func(c *cli.Context) error {
				args := c.Args()
				for i := 0; i < args.Len(); i++ {
					fmt.Println("removed task template: ", c.Args().Get(i))
				}
				return nil
			},
		},
	},
}
