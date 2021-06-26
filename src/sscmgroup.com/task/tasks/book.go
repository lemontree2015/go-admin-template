package tasks

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model"
	"sscmgroup.com/app/model/sys"
)

var BookTasks = &cli.Command{
	Name:  "book",
	Usage: "运行task book 服务",
	Before: func(c *cli.Context) error {
		//注册配制文件
		fmt.Println("book Before")
		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name:  "add",
			Usage: "add a book",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Usage:   "book name",
					Aliases: []string{"n"},
				},
			},
			Action: func(c *cli.Context) (err error) {
				db := model.DbClient("mgr_db")
				u := &sys.User{}
				err = db.Where("status = ? and user_name = ?", 0, c.String("name")).First(u).Error
				logger.Logger.Debug("book add ", c.String("name"))
				fmt.Println("book add ", c.String("name"))
				fmt.Println("find user ", u.UserName)
				return err
			},
		},
		{
			Name:  "del",
			Usage: "del a book",
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
