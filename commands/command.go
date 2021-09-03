package commands

import (
	"github.com/jormin/gacode/entity"
	"github.com/jormin/gacode/helper"
	"github.com/urfave/cli/v2"
)

// Data Save data
var Data *entity.Data

// GA Google Authenticator
var GA *helper.GoogleAuthenticator

// Before
func BeforeFunc(ctx *cli.Context) (err error) {
	Data, err = helper.ReadData()
	GA = helper.NewGoogleAuthenticator()
	return err
}

// After
func AfterFunc(ctx *cli.Context) error {
	return helper.WriteData(Data)
}
