package main

import (
	"log"
	"os"

	_ "github.com/jormin/gacode/commands"
	_ "github.com/jormin/gacode/commands/account"
	"github.com/jormin/gacode/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "gacode",
		Usage:    "The tool to manage accounts and codes of Google Authenticator.",
		Version:  "v1.0.0",
		Commands: config.GetRegisteredCommands(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
