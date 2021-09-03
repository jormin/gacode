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
			Name:      "gen",
			Usage:     "Generate a new account",
			Action:    Generate,
			ArgsUsage: "[name: account name]",
			Before:    commands.BeforeFunc,
			After:     commands.AfterFunc,
		},
	)
}

// Generate Generate a new account
func Generate(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.ErrMissingRequiredArgument
	}
	name := ctx.Args().Get(0)
	if name == "" {
		return errors.ErrMissingRequiredArgument
	}
	for _, v := range commands.Data.Accounts {
		if v.Name == name {
			return errors.ErrAccountNameExists
		}
	}
	secret, err := commands.GA.GenerateSecret()
	if err != nil {
		return errors.ErrGenerateSecret
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
	fmt.Printf("generate account success.\nname: %s\nsecret: %s\nqrcode:%s\n", name, secret, qrCode)
	return nil
}
