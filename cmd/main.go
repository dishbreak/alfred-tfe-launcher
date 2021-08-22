package main

import "github.com/alecthomas/kong"

type Context struct{}

type Cli struct {
	Login  LoginCmd  `cmd:"login" help:"Set TFE User API Key in keyring."`
	Logout LogoutCmd `cmd:"logout" help:"Remove TFE User API Key from keyring."`
}

func main() {
	cli := &Cli{}
	ctx := kong.Parse(cli, kong.Name("tfe-browser"), kong.Description("A CLI tool for use with Alfred"))
	err := ctx.Run(&Context{})
	ctx.FatalIfErrorf(err)
}
