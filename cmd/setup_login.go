package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
)

type LoginCmd struct {
}

func (l *LoginCmd) Help() string {
	return `
Saves a Terraform Cloud or Terraform Enterprise token to the system keychain.
For details on generating a token, check out Terraform documentation.
https://www.terraform.io/docs/cloud/users-teams-organizations/users.html#api-tokens

Note that for the sake of security, tokens must be passed via stdin. For example:

	echo "TOKEN" | tfe-browser setup login

When run outside of a pipe, the command will prompt you for a token on input.
`
}

// Run will save the TFE token to the System Keychain.
func (l *LoginCmd) Run(ctx *Context) error {
	var token string
	resp := lib.NewScriptActionResponse()

	// This command is intended to never return an error. Alfred will read stdout
	// and handle errors within the workflow.
	defer lib.RecoverIfErr(resp)

	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		token = string(b)
	} else {
		fmt.Println("Enter user token. Press <ENTER> when done.")
		reader := bufio.NewReader(os.Stdin)
		token, _ = reader.ReadString('\n')
	}

	token = strings.TrimSpace(token)

	if len(token) == 0 {
		panic(errors.New("cowardly refusing to set an empty token"))
	}

	if err := lib.SetToken(token); err != nil {
		panic(err)
	}
	resp.SetVariable(lib.ExecStatus, lib.StatusOk)
	resp.SendFeedback()
	return nil
}
