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

	echo "TOKEN" | tfe-browser login
`
}

func (l *LoginCmd) Run(ctx *Context) error {
	var token string
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		token = string(b)
	} else {
		fmt.Println("Enter user token. Press <ENTER> when done.")
		reader := bufio.NewReader(os.Stdin)
		token, _ = reader.ReadString('\n')
	}

	token = strings.TrimSpace(token)

	if len(token) == 0 {
		return errors.New("cowardly refusing to set an empty token")
	}
	return lib.SetToken(token)
}
