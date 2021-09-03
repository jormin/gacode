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
			Usage:  "Manage accounts of Google Authenticator",
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}
