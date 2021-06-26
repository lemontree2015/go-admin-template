package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"sscmgroup.com/cmd"
)

var VERSION = "7.0.0"

func main() {
	//必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	ctx := context.Background()

	cliApp := cli.NewApp()
	cliApp.Name = "sscmgroup"
	cliApp.Version = VERSION
	cliApp.Usage = "sscmgroup app"
	cliApp.Commands = []*cli.Command{
		cmd.NewMgrCmd(ctx),
		cmd.NewApiCmd(ctx),
		cmd.NewTaskCmd(ctx),
	}
	//cliApp.Commands = append(cliApp.Commands, )
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
