package account

import (
	"fmt"
	"time"

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
			Name:      "add",
			Usage:     "Add an existing account",
			Action:    Add,
			ArgsUsage: "[name: account name] [secret: account secret]",
			Before:    commands.BeforeFunc,
			After:     commands.AfterFunc,
		},
	)
}

// Add an existing account
func Add(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	name := ctx.Args().Get(0)
	secret := ctx.Args().Get(1)
	if name == "" || secret == "" {
		return errors.MissingRequiredArgumentErr
	}
	for _, v := range commands.Data.Accounts {
		if v.Name == name {
			return errors.AccountNameExistsErr
		}
	}
	qrCode := commands.GA.GetQRCode(name, secret)
	curTime := time.Now().Unix()
	commands.Data.Accounts = append(
		commands.Data.Accounts, &entity.Account{
			Name:       name,
			Secret:     secret,
			QRCode:     qrCode,
			CreateTime: curTime,
			UpdateTime: curTime,
		},
	)
	fmt.Printf("add account success.\nname: %s\nsecret: %s\nqrcode: %s\n", name, secret, qrCode)
	return nil
}
