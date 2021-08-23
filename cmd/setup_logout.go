package main

import (
	"fmt"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
)

type LogoutCmd struct{}

func (l *LogoutCmd) Help() string {
	return "Removes the Terraform token from the system keychain. This cannot be undone!"
}

func (l *LogoutCmd) Run(ctx *Context) error {
	if err := lib.ClearToken(); err != nil {
		fmt.Print("failed -- " + err.Error())
		return nil
	}
	fmt.Print("ok!")
	return nil
}
