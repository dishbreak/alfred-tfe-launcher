package main

import "github.com/dishbreak/terraform-cloud-launcher/lib"

type LogoutCmd struct{}

func (l *LogoutCmd) Help() string {
	return "Removes the Terraform token from the system keychain. This cannot be undone!"
}

func (l *LogoutCmd) Run(ctx *Context) error {
	return lib.ClearToken()
}
