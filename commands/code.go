package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"
	"time"

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
				&cli.BoolFlag{
					Name:        "t",
					Usage:       "Cycle show codes every second",
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
	cycle := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			getAll = ctx.Bool("a")
		case "t":
			cycle = ctx.Bool("t")
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
	if !cycle {
		printCode(accounts)
		return nil
	}
	ticker := time.NewTicker(time.Second)
	for {
		clear()
		printCode(accounts)
		<-ticker.C
	}
}

// printCode print codes of accounts
func printCode(accounts []*entity.Account) {
	contentFormat := "%s\t%s\t%s"
	headers := []interface{}{"Account", "Code", "Remain Time"}
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(w, "%s\n", fmt.Sprintf(contentFormat, headers...))
	for _, item := range accounts {
		code, err := GA.GetCode(item.Secret)
		if err != nil {
			code = "get code error"
		}
		str := fmt.Sprintf(contentFormat, item.Name, code, fmt.Sprintf("%02d Seconds", 30-time.Now().Unix()%30))
		_, _ = fmt.Fprintf(w, "%s\n", str)
	}
	_ = w.Flush()
}

// clear clear screen
func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
