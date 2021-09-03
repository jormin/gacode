package account

import (
	"fmt"

	"github.com/jormin/gacode/commands"
	"github.com/jormin/gacode/config"
	"github.com/jormin/gacode/entity"
	"github.com/jormin/gacode/errors"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"account", &cli.Command{
			Name:      "rm",
			Usage:     "Remove specified or all accounts",
			Action:    Remove,
			ArgsUsage: "[name-1: first account name] [name-2] ... [name-n]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "Remove all accounts",
					Required:    false,
					DefaultText: "false",
				},
			},
			Before: commands.BeforeFunc,
			After:  commands.AfterFunc,
		},
	)
}

// Remove Remove specified or all accounts
func Remove(ctx *cli.Context) error {
	removeAll := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			removeAll = ctx.Bool("a")
		}
	}
	if removeAll {
		commands.Data.Accounts = []*entity.Account{}
		fmt.Println("remove all accounts success")
		return nil
	}
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	for i := 0; i < ctx.Args().Len(); i++ {
		for index, account := range commands.Data.Accounts {
			if account.Name == ctx.Args().Get(i) {
				commands.Data.Accounts = append(commands.Data.Accounts[:index], commands.Data.Accounts[index+1:]...)
				fmt.Printf("remove account %s success\n", account.Name)
			}
		}
	}
	return nil
}
