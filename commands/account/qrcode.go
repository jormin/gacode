package account

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rs/xid"
	"github.com/skip2/go-qrcode"
	"github.com/urfave/cli/v2"

	"github.com/jormin/gacode/commands"
	"github.com/jormin/gacode/config"
	"github.com/jormin/gacode/entity"
	"github.com/jormin/gacode/errors"
	"github.com/jormin/gacode/helper"
)

// init
func init() {
	config.RegisterCommand(
		"account", &cli.Command{
			Name:      "qrcode",
			Aliases:   []string{"qr"},
			Usage:     "Print or Export the QR code image",
			Action:    QRCode,
			ArgsUsage: "[name: account name]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "o",
					Usage:       "Export the QR code to the specified directory",
					Required:    false,
					DefaultText: "",
				},
			},
			Before: commands.BeforeFunc,
			After:  commands.AfterFunc,
		},
	)
}

// QRCode print or export qrcode image
func QRCode(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.ErrMissingRequiredArgument
	}
	name := ctx.Args().Get(0)
	if name == "" {
		return errors.ErrMissingRequiredArgument
	}
	var account *entity.Account
	for _, v := range commands.Data.Accounts {
		if v.Name == name {
			account = v
		}
	}
	if account == nil {
		return errors.ErrAccountNameNotExists
	}
	exportPath := ""
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "o":
			exportPath = ctx.String("o")
			if exportPath == "" {
				return errors.ErrInvalidExportPath
			}
		}
	}
	width := helper.PrintQRCodeSize
	u, _ := user.Current()
	path := u.HomeDir
	if exportPath != "" {
		width = helper.ExportQRCodeSize
		path = exportPath
	}
	file := fmt.Sprintf("%s/%s.png", path, xid.New().String())
	err := qrcode.WriteFile(account.QRCode, qrcode.Highest, width, file)
	if err != nil {
		return err
	}
	if exportPath != "" {
		fmt.Printf("export the QR code success: %s\n", file)
		return nil
	}
	defer os.Remove(file)
	return helper.PrintQRCode(file)
}
