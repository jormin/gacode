package commands

import (
	"github.com/jormin/gacode/config"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:   "account",
			Usage:  "账户信息管理",
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}
