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
			Usage:     "Show codes of specified or all accounts",
			Action:    Code,
			ArgsUsage: "[name-1: first account name] [name-2] ... [name-n]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "Show codes of all accounts",
					Required:    false,
					DefaultText: "false",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Code Show codes of Google Authenticator
func Code(ctx *cli.Context) error {
	getAll := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			getAll = ctx.Bool("a")
		}
	}
	var accounts []*entity.Account
	if getAll {
		accounts = Data.Accounts
	} else {
		if ctx.Args().Len() == 0 {
			return errors.ErrMissingRequiredArgument
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
