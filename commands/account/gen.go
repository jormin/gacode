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
			Usage:     "生成账户信息",
			Action:    Generate,
			ArgsUsage: "[name: 账户名称]",
			Before:    commands.BeforeFunc,
			After:     commands.AfterFunc,
		},
	)
}

// Generate 生成账户信息
func Generate(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	name := ctx.Args().Get(0)
	if name == "" {
		return errors.MissingRequiredArgumentErr
	}
	for _, v := range commands.Data.Accounts {
		if v.Name == name {
			return errors.AccountNameExistsErr
		}
	}
	secret, err := commands.GA.GenerateSecret()
	if err != nil {
		return errors.GenerateSecretErr
	}
	qrCode := commands.GA.GetQRCode(name, secret)
	curTime := time.Now().Unix()
	commands.Data.Accounts = append(
		commands.Data.Accounts, entity.Account{
			Name:       name,
			Secret:     secret,
			QRCode:     qrCode,
			CreateTime: curTime,
			UpdateTime: curTime,
		},
	)
	fmt.Println(fmt.Sprintf("生成账户成功，二维码: %s", qrCode))
	return nil
}
