package main

import (
	"context"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
	"github.com/gen2brain/beeep"
)

type ValidateCmd struct{}

func (v *ValidateCmd) Help() string {
	return `
Checks the currently logged in account and verifies its ability to access TFE API. 
If successful, emits a notification to macOS with the name of the user. 
`
}

func (v *ValidateCmd) Run(ctx *Context) error {
	client, err := lib.NewTfeClient()
	if err != nil {
		return err
	}

	user, err := client.Users.ReadCurrent(context.Background())
	if err != nil {
		beeep.Alert("Failed to authenticate to TFE", err.Error(), "")
		return err
	}

	return beeep.Notify("Connected to TFE", user.Email+" is a valid user!", "")
}
