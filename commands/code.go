package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jormin/gacode/config"
	"github.com/jormin/gacode/entity"
	"github.com/jormin/gacode/errors"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "code",
			Usage:     "获取账户验证码",
			Action:    Code,
			ArgsUsage: "[name-1: 第一个账户名称] [name-2] ... [name-n]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "获取所有账户的验证码",
					Required:    false,
					DefaultText: "false",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Code 获取账户验证码
func Code(ctx *cli.Context) error {
	getAll := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			getAll = ctx.Bool("a")
		}
	}
	var accounts []entity.Account
	if getAll {
		accounts = Data.Accounts
	} else {
		if ctx.Args().Len() == 0 {
			return errors.MissingRequiredArgumentErr
		}
		for i := 0; i < ctx.Args().Len(); i++ {
			for _, account := range Data.Accounts {
				if account.Name == ctx.Args().Get(i) {
					accounts = append(accounts, account)
				}
			}
		}
	}
	contentFormat := "%s\t%s"
	headers := []interface{}{"Account", "Code"}
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(w, "%s\n", fmt.Sprintf(contentFormat, headers...))
	for _, item := range accounts {
		code, err := GA.GetCode(item.Secret)
		if err != nil {
			code = "get code error"
		}
		str := fmt.Sprintf(contentFormat, item.Name, code)
		_, _ = fmt.Fprintf(w, "%s\n", str)
	}
	_ = w.Flush()
	return nil
}
