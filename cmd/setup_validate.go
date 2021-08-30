package main

import (
	"context"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
)

type ValidateCmd struct{}

// Help will display additional help text for viewing at the command line.
func (v *ValidateCmd) Help() string {
	return `
Checks the currently logged in account and verifies its ability to access TFE API. 
If successful, emits a notification to macOS with the name of the user. 
`
}

// Run will attempt to construct a client using the saved token.
func (v *ValidateCmd) Run(ctx *Context) error {
	resp := lib.NewScriptActionResponse()

	// This function is intended to return nil no matter what, because Alfred will
	// handle errors downstream.
	defer resp.RecoverIfErr()

	client, err := lib.NewTfeClient()
	if err != nil {
		panic(err)
	}

	user, err := client.Users.ReadCurrent(context.Background())
	if err != nil {
		panic(err)
	}

	resp.SetVariable(lib.ExecStatus, lib.StatusOk)
	resp.SetVariable("user", user.Email)
	resp.SendFeedback()
	return nil
}
