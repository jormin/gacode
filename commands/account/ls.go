package account

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jormin/gacode/commands"
	"github.com/jormin/gacode/config"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"account", &cli.Command{
			Name:   "ls",
			Usage:  "List all accounts configured",
			Action: List,
			Before: commands.BeforeFunc,
			After:  commands.AfterFunc,
		},
	)
}

// List List all accounts configured
func List(ctx *cli.Context) error {
	accounts := commands.Data.Accounts
	contentFormat := "%s\t%s\t%s"
	headers := []interface{}{"Name", "Secret", "QRCode"}
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(w, "%s\n", fmt.Sprintf(contentFormat, headers...))
	for _, item := range accounts {
		str := fmt.Sprintf(contentFormat, item.Name, item.Secret, item.QRCode)
		_, _ = fmt.Fprintf(w, "%s\n", str)
	}
	_ = w.Flush()
	return nil
}
