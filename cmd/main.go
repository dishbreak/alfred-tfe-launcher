package main

import "github.com/alecthomas/kong"

type Context struct{}

type SetupCmd struct {
	Login    LoginCmd    `cmd:"login" help:"Set TFE User API Key in keyring."`
	Logout   LogoutCmd   `cmd:"logout" help:"Remove TFE User API Key from keyring."`
	Validate ValidateCmd `cmd:"validate" help:"Test API access."`
}

type WorkspaceCmd struct {
	List WorkspaceListCmd `cmd:"list" help:"List TFE workspaces."`
}

type Cli struct {
	Setup      SetupCmd     `cmd:"setup" help:"Commands for setting up API access."`
	Workspaces WorkspaceCmd `cmd:"workspaces" help:"Commands for working with workspaces"`
}

func main() {
	cli := &Cli{}
	ctx := kong.Parse(cli, kong.Name("tfe-browser"), kong.Description("A CLI tool for use with Alfred"))
	err := ctx.Run(&Context{})
	ctx.FatalIfErrorf(err)
}
